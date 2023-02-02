package main

import (
	"fmt"
	"infoSfera_proxy/internal/config"
	"infoSfera_proxy/internal/models"
	"infoSfera_proxy/internal/routes"
	"infoSfera_proxy/pkg/save_file"
	"infoSfera_proxy/pkg/send_request"
	"log"
	"net/http"
	"time"
)

func main() {
	err := run()
	if err != nil {
		log.Fatal(err)
	}
	defer close(config.App.SaveFileChan)
	defer close(config.App.SendRequest)

	port := fmt.Sprintf(":%s", config.App.Env.GetString("server.port"))
	server := http.Server{
		Addr:           port,
		Handler:        routes.Route(),
		TLSConfig:      nil,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	err = server.ListenAndServe()
	if err != nil {
		return
	}
}

func run() error {
	err := config.GetEnv()
	if err != nil {
		log.Println(err)
		return err
	}

	saveFineChan := make(chan models.SaveFileData)
	config.App.SaveFileChan = saveFineChan

	sendRequestChan := make(chan models.Credentials)
	config.App.SendRequest = sendRequestChan

	save_file.ListenToSaveFile()
	send_request.ListenerSendRequest()
	return nil
}
