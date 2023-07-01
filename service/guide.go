package service

import (
	"sfts/initial"
	"sfts/model"
	"time"
)

func CreateGuide(guide model.Guide) error {
	guide.CreateTime = time.Now()

	err := initial.Database.Table("guides").Create(&guide).Error
	if err != nil {
		return err
	}
	return nil
}

func GetTitle() ([]model.Guide, error) {
	var guide []model.Guide
	err := initial.Database.Table("guides").Where("id > 0").Order("create_time DESC").Find(&guide).Error
	if err != nil {
		return nil, err
	}
	return guide, nil
}

func GetGuideById(id int64) (model.Guide, error) {
	var guide model.Guide
	err := initial.Database.Table("guides").Where("id = ?", id).Find(&guide).Error
	if err != nil {
		return guide, err
	}
	return guide, nil
}

func GetGuideByAuthor(author string) ([]model.Guide, error) {
	var guide []model.Guide
	err := initial.Database.Table("guides").Where("author = ?", author).Find(&guide).Error
	if err != nil {
		return guide, err
	}
	return guide, nil
}

func ReadCount(id int64) error {
	var guide model.Guide
	err := initial.Database.Table("guides").Where("id = ?", id).Find(&guide).Error
	if err != nil {
		return err
	}

	guide.ReadCount++
	initial.Database.Save(&guide)
	return nil
}

func RemoveGuide(id int64) error {
	err := initial.Database.Where("id = ?", id).Delete(&model.Guide{}).Error
	if err != nil {
		return err
	}
	return nil
}
