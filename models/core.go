package models

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"time"
)

var (
	db *gorm.DB
)

func init() {
	var err error
	driverName := beego.AppConfig.String("mysqldrivername")
	host := beego.AppConfig.String("mysqlhost")
	port := beego.AppConfig.String("mysqlport")
	username := beego.AppConfig.String("mysqlusername")
	password := beego.AppConfig.String("mysqlpassword")
	database := beego.AppConfig.String("mysqldatabase")
	charset := beego.AppConfig.String("mysqlcharset")
	conn := fmt.Sprintf("%s:%s@(%s:%s)/%s?charset=%s&parseTime=True&loc=Local",
		username,
		password,
		host,
		port,
		database,
		charset)
	db, err = gorm.Open(driverName, conn)
	if err != nil {
		panic(err)
	}
	db.AutoMigrate(new(User), new(Note), new(Message), new(PraiseLog))
	//db.LogMode(true)
	//defer db.Close()
}

type Model struct {
	ID        uint       `gorm:"primary_key" json:"id"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeletedAt *time.Time `sql:"index" json:"deleted_at"`
}
