[English](README.md) | 中文 

# LManusGo
LManusGo 是一个基于 [langchaingo](https://github.com/tmc/langchaingo) 构建的智能体开源项目，基于 [OpenManus](https://github.com/FoundationAgents/OpenManus) 项目的核心设计理念。

无需复杂的环境依赖，仅需简单配置即可快速上手。

## 安装指南
使用LManusGo需要先安装[Chrome](https://support.google.com/chrome/answer/95346?hl=zh-Hans&co=GENIE.Platform%3DDesktop)浏览器，`WebSearch`功能需要使用到Chrome浏览器


## 配置说明
LManusGo 需要配置使用的 LLM API，请按以下步骤设置：

编辑 config/config.ini 添加 API 密钥和自定义设置
```ini
# 全局 LLM 配置
[llm]
# 要使用的 LLM 模型
model = "qwen-plus-latest"
# API URL
base_url = "https://dashscope.aliyuncs.com/compatible-mode/v1"
# 您的 API 密钥
api_key = "你的APIKEY"
# 控制随机性
temperature = 0.8

# 基本配置
[base]
# 保存文件的路径
save_path = "./data"
```

