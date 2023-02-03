package config

import (
	"github.com/spf13/viper"
	"path/filepath"
)

var appConfig = viper.New()

var path = filepath.Join("resource")

func Init() error {
	appConfig.SetConfigName("app")
	appConfig.SetConfigType("yaml")
	appConfig.SetDefault("server.addr", ":8080")
	appConfig.AddConfigPath(path)
	err := appConfig.SafeWriteConfig()
	return err
}

func GetCfg(key string) string {
	return appConfig.GetString(key)
}
