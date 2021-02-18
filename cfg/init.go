package cfg

import (
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

const path = "conf"
const confPath = path + string(os.PathSeparator) + "app.conf"

const defaultConf = `server.addr=:8888
log.enable=true
log.path=logs
`

var configMap = make(map[string]string)

func InitCfg() error {
	_, err := os.Stat(path)
	if err != nil {
		if os.IsNotExist(err) {
			os.MkdirAll(path, os.ModePerm)
		} else {
			return err
		}
	}
	checkOrCreate(confPath, []byte(defaultConf))
	byte, err := ioutil.ReadFile(confPath)
	if err != nil {
		return err
	}
	configStr := string(byte)
	configKvs := strings.Split(configStr, "\n")
	for _, i2 := range configKvs {
		if i2 != "" {
			configKv := strings.Split(i2, "=")
			if len(configKv) > 1 {
				configMap[configKv[0]] = configKv[1]
			}
		}
	}
	return err
}

func checkOrCreate(path string, byte []byte) error {
	f, err := os.Stat(path)
	if err != nil {
		if os.IsNotExist(err) {
			return ioutil.WriteFile(path, byte, os.ModePerm)
		}
		return err
	}
	if f.IsDir() {
		return errors.New(fmt.Sprintf("初始化失败，勿在当前目录下存放名称为%s的目录", path))
	}
	return nil
}

func GetCfg(key string, defaultVal string) string {
	val, ok := configMap[key]
	if !ok {
		return defaultVal
	}
	return val
}

func GetBoolCfg(key string) bool {
	val, ok := configMap[key]
	if ok {
		bVal, err := strconv.ParseBool(val)
		if err != nil {
			return false
		}
		return bVal
	}
	return false
}
