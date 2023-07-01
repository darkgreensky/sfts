package initial

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

const (
	Host     = "101.43.39.61"
	Port     = "3307"
	Username = "root"
	Password = "sjpsjp"
	DBName   = "Tourism_System"
	Timeout  = "10s"
)

var Database *gorm.DB

func InitMySQL() error {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local&timeout=%s",
		Username, Password, Host, Port, DBName, Timeout)
	var err error
	Database, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	_ = Database
	return err
}
