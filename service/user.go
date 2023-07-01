package service

import (
	"crypto/md5"
	"fmt"
	"io"
	"sfts/initial"
	"sfts/model"
)

func GetMd5(str string) (string, error) {
	m := md5.New()
	_, err := io.WriteString(m, str)
	if err != nil {
		return "", err
	}
	arr := m.Sum(nil)
	return fmt.Sprintf("%x", arr), nil
}

func Register(userName string, password string) (int, error) {
	var sameName int64
	err := initial.Database.Table("users").Where("user_name = ?", userName).Count(&sameName).Error
	if err != nil {
		return -2, err
	}
	if sameName != 0 {
		return -1, nil
	}

	hash, err := GetMd5(password)
	if err != nil {
		return -2, err
	}

	user := model.User{
		UserName: userName,
		Password: hash,
        Money:    100,
	}

	err = initial.Database.Create(&user).Error
	if err != nil {
		return -2, err
	}
	return 0, nil
}

func Login(userName string, password string) (int, error) {
	var user model.User
	result := initial.Database.Table("users").Where("user_name = ?", userName).Find(&user)

	if result.Error != nil {
		return -1, result.Error
	}
	if result.RowsAffected == 0 {
		return -2, nil
	}

	hash, err := GetMd5(password)
	if err != nil {
		return -1, err
	}

	if hash != user.Password {
		return -3, nil
	}

	return 0, nil
}
