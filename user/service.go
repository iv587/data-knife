package user

import (
	"dk/db"
	"errors"
	"gorm.io/gorm"
)

// GetUserByNameAndPass 根据用户密码查询用户
func GetUserByNameAndPass(userName, password string) (*User, error) {
	var user User
	err := db.Executor.Model(&User{}).Where(&User{UserName: userName}).First(&user).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("用户未找到")
		} else {
			return nil, err
		}
	}
	if user.Password != password {
		return nil, errors.New("密码错误")
	}
	return &user, nil
}

func CountAll() (int64, error) {
	var count int64
	err := db.Executor.Model(&User{}).Count(&count).Error
	if err != nil {
		return 0, err
	}
	return count, nil
}

func UpdatePasswd(oldPassword, newPassword string, id int) error {
	var user User
	err := db.Executor.Model(&User{}).First(&user, id).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return errors.New("用户未找到")
		}
		return err
	}
	if user.Password != oldPassword {
		return errors.New("旧密码错误")
	}
	user = User{Id: id, Password: newPassword}
	return db.Executor.Model(&user).Update("password", newPassword).Error
}

func InitUser(user User) error {
	return db.Executor.Transaction(func(tx *gorm.DB) error {
		var count int64
		err := tx.Model(&User{}).Count(&count).Error
		if err != nil {
			return err
		}
		if count > 0 {
			return errors.New("请勿重复安装产品")
		}
		return tx.Model(&User{}).Create(&user).Error
	})
}
