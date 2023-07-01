package service

import (
	"regexp"
	"sfts/initial"
	"sfts/model"
)

func GetInfoById(id int64) (model.Information, error) {
	var data model.Information

	err := initial.Database.Table("informations").Where("id = ?", id).Find(&data).Error
	if err != nil {
		return data, err
	}
	return data, nil
}

func GetAllInfo() ([]model.Information, error) {
	var data []model.Information
	err := initial.Database.Table("informations").Where("id > 0").Find(&data).Error
	if err != nil {
		return nil, err
	}
	return data, nil
}

func SearchInfo(text string) ([]model.Information, error) {
	info, err := GetAllInfo()
	if err != nil {
		return nil, err
	}

	reg := regexp.MustCompile(text)
	data := []model.Information{}

	for _, i := range info {
		if reg.MatchString(i.Title) {
			data = append(data, i)
		}
	}
	return data, nil
}

func GetTicketInfo(user string) ([]model.Ticket, error) {
	var data []model.Ticket

	err := initial.Database.Table("tickets").Where("user_name = ?", user).Order("time DESC").Find(&data).Error
	if err != nil {
		return data, err
	}
	return data, nil
}
