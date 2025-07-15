package search

import (
	"github.com/PuerkitoBio/goquery"
	"github.com/chromedp/chromedp"
	"net/url"
	"strings"
	"time"
)

var _ Search = &BaiduSearch{}

type BaiduSearch struct{}

func (e *BaiduSearch) Call(query string) (items []SearchItem, err error) {
	ctx, cancel := newWork()
	defer cancel()

	var divHTMLs []string
	err = chromedp.Run(ctx,
		chromedp.Navigate(baiduSearchURL+url.QueryEscape(query)),
		chromedp.WaitVisible("div#content_left div.result"),
		chromedp.WaitReady("div#content_left div.result"),

		chromedp.Sleep(5*time.Second),

		chromedp.Evaluate(`
        Array.from(document.querySelectorAll('div#content_left div.result'))
            .map(elm => elm.outerHTML) 
    `, &divHTMLs),
	)
	if err != nil {
		return nil, err
	}

	items = make([]SearchItem, len(divHTMLs))
	idx := 0

	for _, liHTML := range divHTMLs {
		doc, err := goquery.NewDocumentFromReader(strings.NewReader(liHTML))
		if err != nil {
			continue
		}

		title := doc.Find("span.tts-b-hl").Text()
		if title == "" {
			continue
		}

		href, exists := doc.Find("a.sc-link").Attr("href")
		if !exists {
			continue
		}

		description := doc.Find("span.summary-text_560AW").Text()
		if description == "" {
			continue
		}

		items[idx].Title = title
		items[idx].URL = href
		items[idx].Description = description
		idx++
	}

	return items, nil
}
