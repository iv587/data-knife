package db

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.io/iv587/goredis-admin/ledis"
	"io/ioutil"
	"os"
	"sync"
)

type Connection struct {
	Id       int    `json:"id" form:"id"`
	Name     string `json:"name" form:"name"`
	Addr     string `json:"addr" form:"addr"`
	Password string `json:"password" form:"password"`
}

type ConnectionWrap struct {
	Connection
	sync.RWMutex
	client map[int]ledis.Cli
}

var connectionDao = struct {
	sync.RWMutex
	connections map[int]ConnectionWrap
}{}
var redisDbPath = path + string(os.PathSeparator) + "redis.db"

func initConnection() error {
	err := checkOrCreate(redisDbPath, []byte("{}"))
	if err != nil {
		return err
	}
	byte, err := ioutil.ReadFile(redisDbPath)
	if err != nil {
		return err
	}
	err = json.Unmarshal(byte, &connectionDao.connections)
	if err != nil {
		return err
	}
	return nil
}

func UpdateConnection(conn Connection, upPass int) error {
	connectionDao.Lock()
	defer connectionDao.Unlock()
	var newId = 0
	for i, v := range connectionDao.connections {
		if v.Addr == conn.Addr && i != conn.Id {
			return errors.New(fmt.Sprintf("该%s已存在", conn.Addr))
		}
		if i > newId {
			newId = i
		}
	}
	newId = newId + 1
	if conn.Id <= 0 {
		conn.Id = newId
	}
	connWrap, ok := connectionDao.connections[conn.Id]
	if ok {
		if upPass == 1 {
			connWrap.Password = conn.Password
		}
	} else {
		connWrap = ConnectionWrap{
			client: make(map[int]ledis.Cli),
		}
		connWrap.Password = conn.Password
	}
	connWrap.Id = conn.Id
	connWrap.Name = conn.Name
	connWrap.Addr = conn.Addr
	connectionDao.connections[conn.Id] = connWrap
	bytes, err := json.Marshal(connectionDao.connections)
	if err != nil {
		return err
	}
	return ioutil.WriteFile(redisDbPath, bytes, os.ModePerm)
}

func GetConnectionById(id int) (Connection, error) {
	connectionDao.RLock()
	defer connectionDao.RUnlock()
	conn, isExist := connectionDao.connections[id]
	if !isExist {
		return Connection{}, errors.New("该实例不存在")
	}
	return Connection{
		Addr:     conn.Addr,
		Password: conn.Password,
		Name:     conn.Name,
	}, nil
}

// 获取连接列表
func ConnectionList() ([]Connection, error) {
	connectionDao.RLock()
	defer connectionDao.RUnlock()
	list := make([]Connection, 0, len(connectionDao.connections))
	for _, i2 := range connectionDao.connections {
		conn := Connection{
			Addr: i2.Addr,
			Name: i2.Name,
			Id:   i2.Id,
		}
		list = append(list, conn)
	}
	return list, nil
}

func Delete(id int) error {
	connectionDao.Lock()
	defer connectionDao.Unlock()
	delete(connectionDao.connections, id)
	bytes, err := json.Marshal(connectionDao.connections)
	if err != nil {
		return err
	}
	return ioutil.WriteFile(redisDbPath, bytes, os.ModePerm)
}

// 获取指定 dbNo的 客户端
func GetClient(id, dbNo int) (ledis.Cli, error) {
	conn, isExist := connectionDao.connections[id]
	if !isExist {
		return nil, errors.New("连接redis失败")
	}
	return ledis.Create(conn.Addr, conn.Password, dbNo)
}
