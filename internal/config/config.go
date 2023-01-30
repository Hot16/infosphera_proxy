package config

import (
	"github.com/spf13/viper"
	"infoSfera_proxy/internal/models"
)

type AppConfig struct {
	Env          *viper.Viper
	SaveFileChan chan models.SaveFileData
}

var App AppConfig

func GetEnv() error {
	viper.SetConfigFile(".env")
	viper.SetConfigType("json")
	if err := viper.ReadInConfig(); err != nil {
		return err
	}
	App.Env = viper.GetViper()
	return nil
}
