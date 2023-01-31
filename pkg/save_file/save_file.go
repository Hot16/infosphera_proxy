package save_file

import (
	"fmt"
	"infoSfera_proxy/internal/config"
	"infoSfera_proxy/internal/models"
	"log"
	"os"
)

func ListenToSaveFile() {
	go func() {
		for {
			fileData := <-config.App.SaveFileChan
			go saveFile(&fileData)
		}
	}()
}

func saveFile(s *models.SaveFileData) {
	log.Println("Save file: ", s.FileName, " start")
	filePath, err := getFilePath(s)
	if err != nil {
		log.Println(err)
	}
	file, err := os.Create(filePath)
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			log.Println(err)
		}
	}(file)
	if err != nil {
		log.Println(err)
	}
	_, err = file.WriteString(s.StringData)
	if err != nil {
		log.Println(err)
	}
	log.Println("Save file: ", s.FileName, " end")
}

func getFilePath(s *models.SaveFileData) (string, error) {
	var filePath string
	pathForRequest := config.App.Env.GetString("file_save.request_xml")
	pathForResponse := config.App.Env.GetString("file_save.response_xml")
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
