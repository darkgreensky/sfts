package service

import (
	"sfts/initial"
	"sfts/model"
	"time"
)

func QueryCurrency(user string) (int64, error) {
	var info model.User
	err := initial.Database.Table("users").Where("user_name = ?", user).Find(&info).Error
	if err != nil {
		return -1, err
	}
	return info.Money, nil
}

func Recharge(user string, momey int64) error {
	var info model.User
	err := initial.Database.Table("users").Where("user_name = ?", user).Find(&info).Error
	if err != nil {
		return err
	}

	info.Money += momey
	err = initial.Database.Table("users").Where("user_name = ?", info.UserName).Save(&info).Error
	if err != nil {
		return err
	}
	return nil
}

func BuyTicket(userName string, id int64) (bool, int64, int64, error) {
	var user model.User
	var info model.Information
	err := initial.Database.Table("users").Where("user_name = ?", userName).Find(&user).Error
	if err != nil {
		return false, 0, 0, err
	}

	err = initial.Database.Table("informations").Where("id = ?", id).Find(&info).Error
	if err != nil {
		return false, 0, 0, err
	}

	if int64(info.Price) > user.Money {
		return false, 0, 0, nil
	}

	ticket := model.Ticket{
		UserName: userName,
		Place:    info.Title,
		Cost:     int64(info.Price),
		Time:     time.Now(),
	}

	err = initial.Database.Table("tickets").Create(&ticket).Error
	if err != nil {
		return false, 0, 0, err
	}

	user.Money -= int64(info.Price)
	err = initial.Database.Table("users").Where("user_name = ?", user.UserName).Save(&user).Error
	if err != nil {
		return false, 0, 0, err
	}
	return true, user.Money, int64(info.Price), err
}
