package model

import "time"

type Response struct {
	StatusCode int32  `json:"status_code"`
	StatusMsg  string `json:"status_msg,omitempty"`
	Result     string `json:"result,omitempty"`
}

type User struct {
	id       int64
	UserName string
	Password string
	Money    int64
}

type UserResponse struct {
	StatusCode int32  `json:"status_code"`
	StatusMsg  string `json:"status_msg,omitempty"`
	Token      string `json:"token,omitempty"`
	Result     string `json:"result,omitempty"`
}

type Guide struct {
	ID         int64     `json:"id"`
	Author     string    `json:"author"`
	Title      string    `json:"title"`
	Content    string    `json:"content"`
	CreateTime time.Time `json:"time"`
	ReadCount  int64     `json:"quantity"`
}

type ImageResponse struct {
	Response
	ImageURL string `json:"image_url"`
}

type GetGuides struct {
	Response
	Guides []Guide `json:"guide"`
}

type GetGuide struct {
	Response
	Guides Guide `json:"guide"`
}

type Information struct {
	Id           int64   `json:"id,omitempty"`
	Content      string  `json:"content"`
	Title        string  `json:"title,omitempty"`
	Image        string  `json:"image,omitempty"`
	Ticket       int64   `json:"ticket,omitempty"`
	Lng          float64 `json:"lng,omitempty"`
	Lat          float64 `json:"lat,omitempty"`
	Introduction string  `json:"introduction,omitempty"`
	Opentime     string  `json:"open_time,omitempty"`
	Price        float64 `json:"price,omitempty"`
	Stuprice     float64 `json:"stu_price,omitempty"`
}

type GetInformation struct {
	Response
	Infor Information `json:"infor"`
}

type GetInformations struct {
	Response
	Infor []Information `json:"infor"`
}

type Team struct {
	Id           int64  `json:"id"`
	Leader       string `json:"leader"`
	Title        string `json:"title"`
	Introduction string `json:"introduction"`
	EndTime      string `json:"time"`
	Count        int64  `json:"count"`
}

type GetTeams struct {
	Response
	Teams []Team `json:"team"`
}

type GetTeam struct {
	Response
	Teams   Team     `json:"team"`
	Members []string `json:"members"`
}

type CheckTeam struct {
	Response
	Check bool `json:"check"`
}

type Apply struct {
	Id       int64  `json:"id"`
	UserName string `json:"user"`
	TeamId   int64  `json:"team_id"`
}

type Program struct {
	ID      int64  `json:"id"`
	Content string `json:"content"`
	Title   string `json:"title"`
	Time    string `json:"time"`
	Locate  string `json:"locate"`
	Count   int64  `json:"count"`
}

type GetPrograms struct {
	Response
	Programs []Program `json:"program"`
}

type GetProgram struct {
	Response
	Programs Program `json:"program"`
}

type Comment struct {
	Id       int64     `json:"id"`
	TeamId   int64     `json:"team_id"`
	UserName string    `json:"user_name"`
	Content  string    `json:"content"`
	Time     time.Time `json:"time"`
}

type GetComments struct {
	Response
	Comments []Comment `json:"comments"`
}

type Ticket struct {
	Id       int64     `json:"id"`
	UserName string    `json:"user_name"`
	Place    string    `json:"place"`
	Cost     int64     `json:"cost"`
	Time     time.Time `json:"time"`
}

type GetMember struct {
	Response
	Members []string `json:"members"`
}

type GetMoney struct {
	Response
	Money  int64    `json:"money"`
	Ticket []Ticket `json:"ticket"`
}

type Buy struct {
	Response
	Money int64 `json:"money"`
	Cost  int64 `json:"cost"`
}
