package search

import (
	"context"
	"github.com/chromedp/chromedp"
	"log"
)

var SearchEngine Search

type Search interface {
	Call(query string) (items []SearchItem, err error)
}

func NewBingSearchEngine() {
	SearchEngine = new(BingSearch)
}

func NewBaiduSearchEngine() {
	SearchEngine = new(BaiduSearch)
}

const (
	bingSearchURL  = "https://cn.bing.com/search?q="
	baiduSearchURL = "https://www.baidu.com/s?wd="
)

type SearchItem struct {
	Title       string `json:"title"`
	URL         string `json:"url"`
	Description string `json:"description"`
}

func newWork() (context.Context, context.CancelFunc) {
	// 设置执行选项
	opts := append(chromedp.DefaultExecAllocatorOptions[:],
		chromedp.NoDefaultBrowserCheck,
		chromedp.Flag("headless", false), // 以有头模式运行
		chromedp.Flag("ignore-certificate-errors", true),
		chromedp.Flag("disable-web-security", true),
		chromedp.NoFirstRun,
		chromedp.Flag("mute-audio", false), // 开启声音
		chromedp.UserAgent("Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/91.0.4472.164 Safari/537.36"),
	)

	// 创建执行分配器
	allocCtx, _ := chromedp.NewExecAllocator(context.Background(), opts...)
	ctx, cancel := chromedp.NewContext(allocCtx, chromedp.WithLogf(log.Printf))
	return ctx, cancel
}
