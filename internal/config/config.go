package config

import (
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
	"infoSfera_proxy/internal/models"
)

type AppConfig struct {
	Env          *viper.Viper
	SaveFileChan chan models.SaveFileData
	SendRequest  chan models.Credentials
}

var App AppConfig

func GetEnv() error {
	viper.SetConfigFile(".env")
	viper.SetConfigType("json")
	viper.WatchConfig()
	viper.OnConfigChange(func(e fsnotify.Event) {
		fmt.Println("Config file changed:", e.Name)
	})
	if err := viper.ReadInConfig(); err != nil {
		return err
	}
	App.Env = viper.GetViper()
	return nil
}
