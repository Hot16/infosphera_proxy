package infoshera

import (
	"infoSfera_proxy/internal/config"
	"infoSfera_proxy/internal/models"
	"infoSfera_proxy/pkg/send_request"
	"log"
)

func InfoSpheraRequest(id string, data string) {
	cred := send_request.NewCred("infoshera")
	cred.Id = id
	cred.PostFields = []byte(data)

	response, err := send_request.SendRequest(&cred)
	if err != nil {
		log.Println(err)
	}
	infoSpheraRsponse(&response)

}

func infoSpheraRsponse(r *models.Response) {
	saveFile := models.SaveFileData{
		Id:         r.Id,
		IsRequest:  false,
		FileName:   r.Id,
		StringData: r.Data,
	}
	config.App.SaveFileChan <- saveFile

	cred := send_request.NewCred("vzaimno")
	cred.Id = r.Id
	cred.GetParams["id_request"] = r.Id
	cred.PostFields = []byte(r.Data)

	_, _ = send_request.SendRequest(&cred)
}
