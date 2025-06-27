package config

import (
	"LManusGo/tools/search"
	"github.com/sirupsen/logrus"
	"gopkg.in/ini.v1"
	"os"
	"path/filepath"
	"strings"
)

var Conf = new(Config)

type Config struct {
	LLM  `ini:"llm"`
	Base `ini:"base"`
}

type LLM struct {
	BaseUrl     string  `ini:"base_url"`
	Model       string  `ini:"model"`
	ApiKey      string  `ini:"api_key"`
	MaxTokens   int     `ini:"max_tokens"`
	Temperature float64 `ini:"temperature"`
}

type Base struct {
	SavePath     string `ini:"save_path"`
	SearchEngine string `ini:"search_engine"`
}

// 读取配置文件
func init() {
	err := ini.MapTo(Conf, "config/config.ini")
	if err != nil {
		logrus.Panic("load config failed, err: " + err.Error())
	}

	// LLM
	Conf.LLM.Model = strings.TrimSpace(Conf.LLM.Model)
	if Conf.LLM.Model == "" {
		logrus.Panic("llm model can not be empty")
	}

	Conf.LLM.BaseUrl = strings.TrimSpace(Conf.LLM.BaseUrl)
	if Conf.LLM.BaseUrl == "" {
		logrus.Panic("llm base url can not be empty")
	}

	Conf.LLM.ApiKey = strings.TrimSpace(Conf.LLM.ApiKey)
	if Conf.LLM.ApiKey == "" {
		logrus.Panic("llm api key can not be empty")
	}

	if Conf.LLM.Temperature < 0 || Conf.LLM.Temperature > 2 {
		logrus.Panic("llm temperature must be between 0 and 2")
	}

	// Base
	Conf.Base.SavePath = strings.TrimSpace(Conf.Base.SavePath)
	if Conf.SavePath == "" {
		Conf.SavePath = "./data"
	}
	Conf.SavePath, err = filepath.Abs(Conf.SavePath)
	if err != nil {
		logrus.Panic("create save path failed, err: " + err.Error())
	}
	// 创建单个文件夹
	err = os.MkdirAll(Conf.Base.SavePath, 0755)
	if err != nil {
		logrus.Panic("create save path failed, err: " + err.Error())
	}
	Conf.Base.SearchEngine = strings.TrimSpace(Conf.Base.SearchEngine)
	switch Conf.Base.SearchEngine {
	case "bing":
		search.NewBingSearchEngine()
		return
	case "baidu":
		search.NewBaiduSearchEngine()
		return
	case "":
		return
	default:
		logrus.Panic("search engine not support, ")
	}
}
