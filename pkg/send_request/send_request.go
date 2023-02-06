package send_request

import (
	"bytes"
	"fmt"
	"infoSfera_proxy/internal/config"
	"infoSfera_proxy/internal/models"
	"io"
	"log"
	"net/http"
)

func ListenerSendRequest() {
	go func() {
		for {
			cred := <-config.App.SendRequest
			go sendRequest(&cred)
		}
	}()
}

func NewCred(id string) models.Credentials {

	cred := models.Credentials{
		BaseUrl:   config.App.Env.GetString(fmt.Sprintf("external.%s.baseUrl", id)),
		Method:    config.App.Env.GetString(fmt.Sprintf("external.%s.method", id)),
		Headers:   make(map[string]string),
		GetParams: make(map[string]string),
	}
	for k, v := range config.App.Env.GetStringMapString(fmt.Sprintf("external.%s.headers", id)) {
		cred.Headers[k] = v
	}
	for k, v := range config.App.Env.GetStringMapString(fmt.Sprintf("external.%s.query_params", id)) {
		cred.GetParams[k] = v
	}
	return cred
}

func sendRequest(c *models.Credentials) {
	req, err := http.NewRequest(c.Method, c.BaseUrl, bytes.NewReader(c.PostFields))
	if err != nil {
		log.Println(err)
	}

	queryParams := req.URL.Query()
	for k, v := range c.GetParams {
		queryParams.Add(k, v)
	}

	req.URL.RawQuery = queryParams.Encode()

	for k, v := range c.Headers {
		fmt.Println(k + " " + v)
		req.Header.Add(k, v)
	}

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Println(err)
	}

	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			log.Println(err)
		}
	}(res.Body)
	body, _ := io.ReadAll(res.Body)

	fileData := models.SaveFileData{
		Id:         c.Id,
		IsRequest:  false,
		FileName:   c.Id,
		StringData: string(body),
	}
	config.App.SaveFileChan <- fileData

	cred := NewCred("vzaimno")
	cred.Id = c.Id
	cred.GetParams["id_request"] = c.Id
	c.PostFields = []byte(string(body))
	config.App.SendRequest <- cred
}
