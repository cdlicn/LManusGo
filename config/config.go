package config

import (
	"LManusGo/tools/search"
	"github.com/sirupsen/logrus"
	"gopkg.in/ini.v1"
	"os"
	"path/filepath"
	"runtime"
	"strings"
)

var Conf = new(Config)

type Config struct {
	LLM  `ini:"llm"`
	Base `ini:"base"`
}

type LLM struct {
	ApiType     string  `ini:"api_type"`
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

// 获取项目根目录
func getRootDir() string {
	_, filename, _, _ := runtime.Caller(0)
	return filepath.Dir(filepath.Dir(filename))
}

// 读取配置文件
func init() {
	// 读取配置文件
	configPath := filepath.Join(getRootDir(), "config", "config.ini")
	err := ini.MapTo(Conf, configPath)
	if err != nil {
		logrus.Panic("load config failed, err: " + err.Error())
	}

	// LLM
	Conf.LLM.ApiType = strings.TrimSpace(Conf.LLM.ApiType)
	Conf.LLM.Model = strings.TrimSpace(Conf.LLM.Model)
	Conf.LLM.BaseUrl = strings.TrimSpace(Conf.LLM.BaseUrl)
	Conf.LLM.ApiKey = strings.TrimSpace(Conf.LLM.ApiKey)

	if Conf.LLM.ApiType == "" {
		logrus.Warnln("llm api_type is empty")
		logrus.Warnln("use default api_type openai")
		Conf.LLM.ApiType = "openai"
	}

	if Conf.LLM.Model == "" {
		logrus.Panic("llm model can not be empty")
	}

	if Conf.LLM.BaseUrl == "" {
		logrus.Panic("llm base url can not be empty")
	}

	if Conf.LLM.ApiKey == "" {
		logrus.Panic("llm api key can not be empty")
	}

	if Conf.LLM.Temperature < 0 || Conf.LLM.Temperature > 2 {
		logrus.Warnln("llm temperature must be between 0 and 2")
		logrus.Warnln("use default temperature value 0")
		Conf.LLM.Temperature = 0
	}

	// Base
	Conf.Base.SavePath = strings.TrimSpace(Conf.Base.SavePath)
	if Conf.SavePath == "" {
		logrus.Warnln("save_path is empty")
		logrus.Warnln("use default save path ./data")
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
		logrus.Warnln("search_engine is empty")
		logrus.Warnln("use default search engine bing")
		search.NewBingSearchEngine()
		return
	default:
		logrus.Panic("search engine not support")
	}
}
