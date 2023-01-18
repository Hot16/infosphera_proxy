package save_file

import (
	"fmt"
	"infoSfera_proxy/internal/config"
	"os"
)

type SaveFileData struct {
	IsRequest  bool
	FileName   string
	StringData string
}

func (s *SaveFileData) SaveFile(app *config.AppConfig) error {
	filePath, err := s.getFilePath(app)
	if err != nil {
		return err
	}
	file, err := os.Create(filePath)
	defer file.Close()
	if err != nil {
		return err
	}
	_, err = file.WriteString(s.StringData)
	if err != nil {
		return err
	}

	return nil
}

func (s *SaveFileData) getFilePath(app *config.AppConfig) (string, error) {
	var filePath string
	pathForRequest := app.Env.GetString("file_save.request_xml")
	pathForResponse := app.Env.GetString("file_save.response_xml")
	if s.IsRequest {
		filePath = "." + pathForRequest
	} else {
		filePath = "." + pathForResponse
	}

	//создаёт диреторию, если ее нет
	creatingPath := createDirPath{Path: filePath}
	err := creatingPath.createDir()
	if err != nil {
		return "", err
	}
	filePath = fmt.Sprintf("%s%s.xml", filePath, s.FileName)
	return filePath, nil
}
