package service

import (
	"fmt"
	"sfts/initial"
	"sfts/model"
	"time"
)

func CreateTeam(team model.Team) error {
	err := initial.Database.Table("teams").Create(&team).Error
	if err != nil {
		return err
	}

	return nil
}

func SearchTeam() ([]model.Team, error) {
	var team, team2 []model.Team
	time := fmt.Sprintf("%04d/%02d/%02d", int64(time.Now().Year()), int64(time.Now().Month()), int64(time.Now().Day()))
	err := initial.Database.Table("teams").Where("end_time >= ?", time).Order("end_time").Find(&team).Error
	if err != nil {
		return nil, err
	}

	err = initial.Database.Table("teams").Where("end_time < ?", time).Find(&team2).Error
	if err != nil {
		return nil, err
	}

	for i := 0; i < len(team2); i++ {
		team = append(team, team2[i])
	}

	return team, nil
}

func GetTeamById(id int64) (model.Team, error) {
	var team model.Team
	err := initial.Database.Table("teams").Where("id = ?", id).Find(&team).Error
	if err != nil {
		return team, err
	}
	return team, nil
}

func GetTeamsByUser(user string) ([]model.Team, error) {
	var team1 []model.Team
	var team2 []model.Apply
	err := initial.Database.Table("teams").Where("leader = ?", user).Find(&team1).Error
	if err != nil {
		return nil, err
	}

	err = initial.Database.Table("applies").Where("user_name = ?", user).Find(&team2).Error
	if err != nil {
		return nil, err
	}

	for i := 0; i < len(team2); i++ {
		team, err := GetTeamById(team2[i].TeamId)
		if err != nil {
			return nil, err
		}
		team1 = append(team1, team)
	}

	return team1, nil
}

func CheckAddTeam(user string, id int64) (bool, error) {
	var count int64
	err := initial.Database.Table("applies").Where("user_name = ? and team_id = ?", user, id).Count(&count).Error
	if err != nil {
		return true, err
	}
	if count > 0 {
		return true, nil
	}

	err = initial.Database.Table("teams").Where("id = ? and leader = ?", id, user).Count(&count).Error
	if err != nil {
		return true, err
	}

	return count > 0, nil
}

func AddToTeam(apply model.Apply) error {
	check, err := CheckAddTeam(apply.UserName, apply.TeamId)
	if err != nil {
		return err
	}
	if check {
		return nil
	}

	err = initial.Database.Table("applies").Create(&apply).Error
	if err != nil {
		return err
	}

	var team model.Team
	err = initial.Database.Table("teams").Where("id = ?", apply.TeamId).Find(&team).Error
	if err != nil {
		return err
	}

	team.Count++
	initial.Database.Save(&team)

	return nil
}

func RemoveFromTeam(user string, id int64) error {
	err := initial.Database.Table("applies").Where("user_name = ? and team_id = ?", user, id).Delete(&model.Guide{}).Error
	if err != nil {
		return err
	}

	var team model.Team
	err = initial.Database.Table("teams").Where("id = ?", id).Find(&team).Error
	if err != nil {
		return err
	}

	if team.Count == 1 {
		initial.Database.Delete(&team)
	} else {
		team.Count--
		initial.Database.Save(&team)
	}

	return nil
}

func GetTeamMember(id int64) ([]string, error) {
	var member []string
	var team model.Team
	var apply []model.Apply

	err := initial.Database.Table("teams").Where("id = ?", id).Find(&team).Error
	if err != nil {
		return nil, err
	}

	err = initial.Database.Table("applies").Where("team_id = ?", id).Find(&apply).Error
	if err != nil {
		return nil, err
	}

	member = append(member, team.Leader)
	for i := 0; i < len(apply); i++ {
		member = append(member, apply[i].UserName)
	}
	return member, nil
}
