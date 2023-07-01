package service

import (
	"sfts/initial"
	"sfts/model"
	"time"
)

func CreateComment(comment model.Comment) error {
	comment.Time = time.Now()
	err := initial.Database.Table("comments").Create(&comment).Error
	if err != nil {
		return err
	}
	return nil
}

func GetComment(TeamId int64) ([]model.Comment, error) {
	var comment []model.Comment
	err := initial.Database.Table("comments").Where("team_id = ?", TeamId).Find(&comment).Error
	if err != nil {
		return nil, err
	}
	return comment, nil
}
