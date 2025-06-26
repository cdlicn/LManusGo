package config

import (
	"gopkg.in/ini.v1"
	"os"
	"path/filepath"
)

var Conf = new(Config)

// 读取配置文件
func init() {
	err := ini.MapTo(Conf, "config/config.ini")
	if err != nil {
		panic("load config failed, err: " + err.Error())
	}
	if Conf.SavePath == "" {
		Conf.SavePath = "./data"
	}
	Conf.SavePath, err = filepath.Abs(Conf.SavePath)
	if err != nil {
		panic("create save path failed, err: " + err.Error())
	}
	// 创建单个文件夹
	err = os.MkdirAll(Conf.SavePath, 0755)
	if err != nil {
		panic("create save path failed, err: " + err.Error())
	}
}

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
	SavePath string `ini:"save_path"`
}
