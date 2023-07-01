package service

import (
	"sfts/initial"
	"sfts/model"
)

func SearchCount(id int64) error {
	var program model.Program
	err := initial.Database.Table("programs").Where("id = ?", id).Find(&program).Error
	if err != nil {
		return err
	}

	program.Count++
	initial.Database.Save(&program)
	return nil
}

func GetProgramById(id int64) (model.Program, error) {
	var data model.Program
	err := initial.Database.Table("programs").Where("id = ?", id).Find(&data).Error
	if err != nil {
		return data, err
	}

	err = SearchCount(id)
	if err != nil {
		return data, err
	}

	return data, nil
}

func GetProgramByTime() ([]model.Program, error) {
	var data []model.Program
	err := initial.Database.Table("programs").Where("id > 0").Order("start_time").Find(&data).Error
	if err != nil {
		return nil, err
	}
	return data, nil
}

func GetProgramByCount() ([]model.Program, error) {
	var data []model.Program
	err := initial.Database.Table("programs").Where("id > 0").Order("count DESC").Find(&data).Error
	if err != nil {
		return nil, err
	}
	return data, nil
}
