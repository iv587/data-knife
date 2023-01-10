package connection

import (
	"dk/db"
	"errors"
	"gorm.io/gorm"
	"strings"
)

type Connection struct {
	Id       int    `json:"id" gorm:"primaryKey" form:"id"`
	Name     string `json:"name" form:"name"`
	Addr     string `json:"addr" form:"addr"`
	Password string `json:"password" form:"password"`
}

func (Connection) TableName() string {
	return "connection"
}

// List 获取列表
func List() ([]Connection, error) {
	var list []Connection
	if err := db.Executor.Model(&Connection{}).Find(&list).Error; err != nil {
		return nil, err
	}
	return list, nil
}

// GetById 获取详情
func GetById(id int) (*Connection, error) {
	var connection Connection
	if err := db.Executor.Model(&Connection{}).First(&connection, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("该记录不存在！")
		}
		return nil, err
	}
	return &connection, nil
}

func AddOrUpdate(data Connection) error {
	return db.Executor.Transaction(func(tx *gorm.DB) error {
		if data.Id > 0 {
			ex := tx
			if strings.EqualFold(data.Password, "") {
				ex = tx.Omit("password")
			}
			return ex.Save(&data).Error
		} else {
			return tx.Create(&data).Error
		}
	})
}
