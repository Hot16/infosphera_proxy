package save_file

import (
	"fmt"
	"infoSfera_proxy/internal/config"
	"os"
)

type FileData struct {
	IsRequest  bool
	FileName   string
	StringData string
}

func SaveFile(app config.AppConfig, s FileData) error {
	filePath, err := getFilePath(app, s)
	if err != nil {
		return err
	}
	file, err := os.Create(fmt.Sprintf("%s%s.xml", filePath, s.FileName))
	if err != nil {
		return err
	}
	_, err = file.WriteString(s.StringData)
	if err != nil {
		return err
	}

	return nil
}

func getFilePath(app config.AppConfig, s FileData) (string, error) {
	var filePath string
	pathForRequest := app.Env.GetString("file_save.request_xml")
	pathForResponse := app.Env.GetString("file_save.response_xml")
	if s.IsRequest {
		filePath = "." + pathForRequest
	} else {
		filePath = "." + pathForResponse
	}
	//создаёт диреторию, если ее нет
	err := os.MkdirAll(filePath, os.ModePerm)
	if err != nil {
		return "", err
	}
	filePath = fmt.Sprintf("%s%s.xml", filePath, s.FileName)
	return filePath, nil
}
