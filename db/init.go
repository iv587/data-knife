package db

import (
	"errors"
	"fmt"
	"io/ioutil"
	"os"
)

const path = "data"

func Init() error {
	// 查找配置文件并初始化用户
	_, err := os.Stat(path)
	if err != nil {
		if os.IsNotExist(err) {
			os.MkdirAll(path, os.ModePerm)
		} else {
			return err
		}
	}
	err = initUser()
	if err != nil {
		return err
	}
	err = initConnection()
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
