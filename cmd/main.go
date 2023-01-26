package main

import (
	"fmt"
	"infoSfera_proxy/internal/config"
	"infoSfera_proxy/internal/routes"
	"log"
	"net/http"
	"time"
)

func main() {
	err := config.GetEnv()
	if err != nil {
		log.Println(err)
	}
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
