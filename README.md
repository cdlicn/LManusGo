English| [中文](README_zh.md)

# LManusGo
LManusGo is an open-source agent project built on [langchaingo](https://github.com/tmc/langchaingo), following the core design philosophy of the [OpenManus](https://github.com/FoundationAgents/OpenManus) project.

No complex dependencies required—just simple configuration to get started quickly.

## Installation Guide
To use LManusGo, you need to install [Chrome](https://www.google.com/chrome/) first, as the `WebSearch` feature relies on it.

## Configuration
LManusGo requires configuring the LLM API. Follow these steps:

Edit `config/config.ini` to add your API key and customize settings:
```ini  
# Global LLM Configuration  
[llm]
# LLM model to use  
model = "qwen-plus-latest"
# API URL  
base_url = "https://dashscope.aliyuncs.com/compatible-mode/v1"
# Your API key  
api_key = "your_api_key_here"
# Controls randomness  
temperature = 0.8

# Basic Configuration  
[base]
# Path to save files  
save_path = "./data"  
# search engines [baidu, bing]
search_engine = "bing"
```