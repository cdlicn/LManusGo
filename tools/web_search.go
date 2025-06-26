package tools

import (
	"LManusGo/tools/search"
	"context"
	"encoding/json"
	"fmt"
	"github.com/tmc/langchaingo/llms"
)

var WebSearchConfig webSearchConfig

type webSearchConfig struct {
	ApiKey string
	Engine string
}

var _ Tool = &SearchWeb{}

type SearchWeb struct{}

func (t SearchWeb) Name() string {
	return "SearchWeb"
}
func (t SearchWeb) GetTool() llms.Tool {
	return llms.Tool{
		Type: "function",
		Function: &llms.FunctionDefinition{
			Name:        "SearchWeb",
			Description: `Search for information from Baidu search engine`,
			Parameters: map[string]any{
				"type": "object",
				"properties": map[string]any{
					"query": map[string]any{
						"type":        "string",
						"description": "keyword that need to be searched",
					},
				},
				"required": []string{"query"},
			},
		},
	}
}

func (t SearchWeb) Call(ctx context.Context, input string) string {
	var mp map[string]string
	err := json.Unmarshal([]byte(input), &mp)
	if err != nil {
		return "解析输入出错了" + err.Error()
	}

	engine := search.NewSearchEngine("bing")
	items, err := engine.BingSearch(mp["query"])

	return fmt.Sprintln(items)
}

// SearchResult 定义主结构体
type SearchResult struct {
	SearchMetadata   SearchMetadata   `json:"search_metadata"`
	SearchParameters SearchParameters `json:"search_parameters"`
	Ads              []Ad             `json:"ads"`
	OrganicResults   []OrganicResult  `json:"organic_results"`
	TopSearches      []TopSearch      `json:"top_searches"`
	TopStories       []TopStory       `json:"top_stories"`
	RelatedSearches  []RelatedSearch  `json:"related_searches"`
	PeopleAlsoSearch []RelatedSearch  `json:"people_also_search_for"`
	Pagination       Pagination       `json:"pagination"`
}

// SearchMetadata 搜索元数据
type SearchMetadata struct {
	ID               string  `json:"id"`
	Status           string  `json:"status"`
	CreatedAt        string  `json:"created_at"`
	RequestTimeTaken float64 `json:"request_time_taken"`
	ParsingTimeTaken float64 `json:"parsing_time_taken"`
	TotalTimeTaken   float64 `json:"total_time_taken"`
	RequestURL       string  `json:"request_url"`
	HTMLURL          string  `json:"html_url"`
	JSONURL          string  `json:"json_url"`
}

// SearchParameters 搜索参数
type SearchParameters struct {
	Engine string `json:"engine"`
	Query  string `json:"q"`
}

// Ad 广告
type Ad struct {
	Position  int        `json:"position"`
	Title     string     `json:"title"`
	Link      string     `json:"link"`
	Source    string     `json:"source"`
	Date      string     `json:"date"`
	Snippet   string     `json:"snippet"`
	Thumbnail string     `json:"thumbnail"`
	Sitelinks *Sitelinks `json:"sitelinks,omitempty"`
}

// Sitelinks 站点链接
type Sitelinks struct {
	Expanded []ExpandedSitelink `json:"expanded"`
}

// ExpandedSitelink 扩展站点链接
type ExpandedSitelink struct {
	Title     string `json:"title"`
	Link      string `json:"link"`
	Snippet   string `json:"snippet"`
	Thumbnail string `json:"thumbnail"`
}

// OrganicResult 有机搜索结果
type OrganicResult struct {
	Position                int      `json:"position"`
	Title                   string   `json:"title"`
	Link                    string   `json:"link"`
	DisplayedLink           string   `json:"displayed_link"`
	Snippet                 string   `json:"snippet"`
	SnippetHighlightedWords []string `json:"snippet_highlighted_words"`
	Date                    string   `json:"date"`
	Thumbnail               string   `json:"thumbnail"`
}

// TopSearch 热门搜索
type TopSearch struct {
	Position int    `json:"position"`
	Query    string `json:"query"`
	Link     string `json:"link"`
	IsHot    bool   `json:"is_hot,omitempty"`
}

// TopStory 热门故事
type TopStory struct {
	Title     string `json:"title"`
	Link      string `json:"link"`
	Source    string `json:"source"`
	Date      string `json:"date"`
	Snippet   string `json:"snippet"`
	Thumbnail string `json:"thumbnail"`
}

// RelatedSearch 相关搜索
type RelatedSearch struct {
	Query string `json:"query"`
	Link  string `json:"link"`
}

// Pagination 分页信息
type Pagination struct {
	Current    int               `json:"current"`
	Next       string            `json:"next"`
	OtherPages map[string]string `json:"other_pages"`
}
