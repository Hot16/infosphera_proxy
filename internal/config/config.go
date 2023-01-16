package config

import "github.com/spf13/viper"

type AppConfig struct {
	Env *viper.Viper
}
