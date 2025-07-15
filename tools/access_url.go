package tools

import (
	"context"
	"github.com/chromedp/chromedp"
	"github.com/tmc/langchaingo/llms"
	"log"
	"time"
)

var _ Tool = &AccessURL{}

type AccessURL struct {
}

func (t AccessURL) Name() string {
	return "AccessURL"
}

func (t AccessURL) GetTool() llms.Tool {
	return llms.Tool{
		Type: "function",
		Function: &llms.FunctionDefinition{
			Name:        "AccessURL",
			Description: `use browser to access the url. you'll get the <body> code for the html page'`,
			Parameters: map[string]any{
				"type": "object",
				"properties": map[string]any{
					"url": map[string]any{
						"type":        "string",
						"description": "the url address you want to access, the url must start with 'http' or 'https'",
					},
				},
				"required": []string{"url"},
			},
		},
	}
}
func (t AccessURL) Call(ctx context.Context, input string) string {
	mp, err := unmarshallJson(input)
	if err != nil {
		return err.Error()
	}
	url := mp["url"]

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
	defer cancel()

	var text string

	err = chromedp.Run(ctx,
		chromedp.Navigate(url),
		chromedp.WaitReady("body"),

		chromedp.Sleep(5*time.Second),

		chromedp.Text("body", &text),
	)
	if err != nil {
		return "failed to access url, error: " + err.Error() + ". Please retry"
	}
	return text
}
