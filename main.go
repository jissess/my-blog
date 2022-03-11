package main

import (
	"encoding/gob"
	"github.com/astaxie/beego"
	"my-blog/models"
	_ "my-blog/models"
	_ "my-blog/routers"
	"strings"
)

func main() {
	initSession()
	initTemplate()
	beego.Run()
}

func initSession() {
	gob.Register(models.User{})
	beego.BConfig.WebConfig.Session.SessionOn = true
	beego.BConfig.WebConfig.Session.SessionName = "blog"
	beego.BConfig.WebConfig.Session.SessionProvider = "file"
	beego.BConfig.WebConfig.Session.SessionProviderConfig = "data/session"
}

func initTemplate() {
	beego.AddFuncMap("equrl", func(x, y string) bool {
		x1 := strings.Trim(x, "/")
		y1 := strings.Trim(y, "/")
		return strings.Compare(x1, y1) == 0
	})
	beego.AddFuncMap("add", func(x, y int) int {
		return x + y
	})
}
