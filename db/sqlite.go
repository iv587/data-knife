package db

import (
	"dk/file"
	"errors"
	"fmt"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"os"
	"path/filepath"
	"time"
)

const ddlScript = `create table connection
(
    id       integer
        constraint connection_pk
            primary key autoincrement,
    name     TEXT default '' not null,
    addr     TEXT default '' not null,
    password text default '' not null
);
create table user
(
    id        INTEGER not null
        constraint user_pk
            primary key autoincrement,
    user_name text    not null,
    password  text    not null
);`

var Executor *gorm.DB

func InitConnection() error {
	// 初始化日志加载器
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer（日志输出的目标，前缀和日志包含的内容——译者注）
		logger.Config{
			SlowThreshold:             time.Second, // 慢 SQL 阈值
			LogLevel:                  logger.Info, // 日志级别
			IgnoreRecordNotFoundError: true,        // 忽略ErrRecordNotFound（记录未找到）错误
			Colorful:                  true,        // 禁用彩色打印
		},
	)
	execDdl := false
	var err error
	fileInfo, err := os.Stat(file.DbPath)
	if err != nil {
		if os.IsNotExist(err) {
			err = file.CreatFile(file.DbPath)
			execDdl = true
			if err != nil {
				return err
			}
		} else {
			return err
		}
	} else if fileInfo.IsDir() {
		return errors.New("存在同名的文件路径，数据初始化失败")
	}
	Executor, err = gorm.Open(sqlite.Open(filepath.Join("resource", "data", "data.sqlite")), &gorm.Config{
		Logger: newLogger,
	})
	if err != nil {
		return err
	}
	if execDdl {
		err = Executor.Exec(ddlScript).Error
		if err != nil {
			return err
		}
	}
	return nil
}

func Test() {
	InitConnection()
	fmt.Println(Executor.Name())
}
