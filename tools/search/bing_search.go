package search

import (
	"github.com/PuerkitoBio/goquery"
	"github.com/chromedp/chromedp"
	"net/url"
	"strings"
	"time"
)

func (e *searchEngine) BingSearch(query string) (items []SearchItem, err error) {

	ctx, cancel := newWork()
	defer cancel()

	var liHTMLs []string
	err = chromedp.Run(ctx,
		chromedp.Navigate(bingSearchURL+url.QueryEscape(query)),
		chromedp.WaitVisible("ol#b_results li.b_algo"),
		chromedp.WaitReady("ol#b_results li.b_algo"),

		chromedp.Sleep(5*time.Second),

		chromedp.Evaluate(`
        Array.from(document.querySelectorAll('ol#b_results li.b_algo'))
            .map(li => li.outerHTML) 
    `, &liHTMLs),
	)
	if err != nil {
		return nil, err
	}

	items = make([]SearchItem, len(liHTMLs))
	idx := 0

	for _, liHTML := range liHTMLs {
		doc, err := goquery.NewDocumentFromReader(strings.NewReader(liHTML))
		if err != nil {
			continue
		}

		href, exists := doc.Find("a").Attr("href")
		if !exists {
			continue
		}

		title := doc.Find("h2").Text()
		if title == "" {
			continue
		}

		description := doc.Find("div.b_caption").Text()
		if description == "" {
			continue
		}

		items[idx].Title = title
		items[idx].URL = href
		items[idx].Description = description
		idx++
	}

	return items[:idx], nil
}
