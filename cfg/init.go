package cfg

import (
	"github.com/spf13/viper"
)

const path = "conf"

var configMap = make(map[string]string)

var appConf = viper.New()

func InitCfg() error {
	appConf.SetConfigName("app")
	appConf.SetConfigType("yaml")
	appConf.AddConfigPath(path)
	appConf.SetDefault("server.addr", ":8080")
	appConf.SetDefault("log.enable", true)
	appConf.SetDefault("log.path", "logs")
	appConf.SafeWriteConfig()
	err := appConf.ReadInConfig()
	return err
}

func GetCfg(key string) string {
	return appConf.GetString(key)
}

func GetBoolCfg(key string) bool {
	return appConf.GetBool(key)
}
