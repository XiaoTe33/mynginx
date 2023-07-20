package ttviper

import (
	"github.com/spf13/viper"
	"mynginx/pkg/myLog"
)

var log = myLog.Log

type Config struct {
	Viper *viper.Viper
}

func ReadConfig(dir string, filename string) *Config {
	v := viper.New()
	v.SetConfigType("yaml")
	v.AddConfigPath(dir)
	v.SetConfigName(filename)
	log.Info("reading config")
	err := v.ReadInConfig()
	if err != nil {
		log.Errorf("read config err: %v", err)
		return nil
	}
	log.Info("reading config successfully")
	return &Config{Viper: v}
}
