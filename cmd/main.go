package main

import (
	"fmt"
	"github.com/spf13/viper"
	"infoSfera_proxy/internal/config"
	"infoSfera_proxy/internal/routes"
	"net/http"
	"time"
)

var app config.AppConfig

func main() {
	getEnv()
	port := fmt.Sprintf(":%s", app.Env.GetString("server.port"))
	server := http.Server{
		Addr:           port,
		Handler:        routes.Route(&app),
		TLSConfig:      nil,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	err := server.ListenAndServe()
	if err != nil {
		return
	}
}

func getEnv() {
	viper.SetConfigFile(".env")
	viper.SetConfigType("json")
	if err := viper.ReadInConfig(); err != nil {
		panic(fmt.Errorf("fatal error: %q", err))
	}
	app.Env = viper.GetViper()
}
