package ttviper

import (
	"github.com/spf13/viper"
	"mynginx/pkg/myLog"
)

var log = myLog.Log

type Config struct {
	Viper *viper.Viper
}

func ReadConfig() *Config {
	v := &viper.Viper{}
	v.SetConfigType("yaml")
	v.AddConfigPath("/etc")
	v.SetConfigName("config.yaml")
	log.Info("reading config")
	err := v.ReadInConfig()
	if err != nil {
		log.Errorf("read config err")
		return nil
	}
	log.Info("reading config successfully")
	return &Config{Viper: v}
}
