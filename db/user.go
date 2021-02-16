package db

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"os"
	"sync"
)

var userDao = struct {
	sync.RWMutex
	user map[int]User
}{user: make(map[int]User)}

type User struct {
	Id       int    `json:"id" form:"id"`
	Name     string `json:"name" form:"name"`
	Password string `json:"password" form:"password"`
}

var userDbPath = path + string(os.PathSeparator) + "user.db"

func initUser() error {
	initData := make(map[int]User)
	initData[1] = User{Id: 1, Name: "admin", Password: "goredis"}
	byte, err := json.Marshal(initData)
	if err != nil {
		return err
	}
	err = checkOrCreate(userDbPath, byte)
	if err != nil {
		return err
	}
	byte, err = ioutil.ReadFile(userDbPath)
	if err != nil {
		return err
	}
	err = json.Unmarshal(byte, &userDao.user)
	if err != nil {
		return err
	}
	return nil
}

func GetUserByNameAndPass(name, password string) (User, error) {
	userDao.RLock()
	defer userDao.RUnlock()
	for _, user := range userDao.user {
		if user.Name == name && user.Password == password {
			return user, nil
		}
	}
	return User{}, errors.New("用户未找到")
}

func UpdateUser(user User) error {
	userDao.Lock()
	defer userDao.Unlock()
	userDao.user[user.Id] = user
	byte, err := json.Marshal(userDao.user)
	if err != nil {
		return err
	}
	return ioutil.WriteFile(userDbPath, byte, os.ModePerm)
}

func GetUserInfo(id int) (User, error) {
	userDao.RLock()
	defer userDao.RUnlock()
	u, ok := userDao.user[id]
	if ok {
		return u, nil
	}
	return u, errors.New("用户不存在")
}
