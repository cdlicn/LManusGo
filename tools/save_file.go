package tools

import (
	"LManusGo/config"
	"context"
	"github.com/tmc/langchaingo/llms"
	"os"
)

var _ Tool = &SaveFile{}

type SaveFile struct{}

func (t SaveFile) Name() string {
	return "SaveFile"
}
func (t SaveFile) GetTool() llms.Tool {
	return llms.Tool{
		Type: "function",
		Function: &llms.FunctionDefinition{
			Name:        "SaveFile",
			Description: `Save the file locally`,
			Parameters: map[string]any{
				"type": "object",
				"properties": map[string]any{
					"fileName": map[string]any{
						"type":        "string",
						"description": "the name of the file that needs to be saved",
					},
					"content": map[string]any{
						"type":        "string",
						"description": "the content of the file that needs to be saved",
					},
				},
				"required": []string{"query"},
			},
		},
	}
}

func (t SaveFile) Call(ctx context.Context, input string) string {
	mp, err := unmarshallJson(input)
	if err != nil {
		return err.Error()
	}

	filePath := config.Conf.SavePath + "\\" + mp["fileName"]
	content := mp["content"]

	// 判断文件是否存在
	if _, err := os.Stat(filePath); os.IsExist(err) {
		err = os.Remove(filePath)
		if err != nil {
			return "there are duplicate file names, failed to delete file, try again with a different file name"
		}
	}

	err = os.WriteFile(filePath, []byte(content), 0644)
	if err != nil {
		return "failed to save file" + err.Error()
	}
	return "save file successful, the save path is: " + filePath
}
