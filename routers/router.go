package routers

import (
	"github.com/astaxie/beego"
	"my-blog/controllers"
)

func init() {
	beego.ErrorController(&controllers.ErrorController{})
	beego.Router("/", &controllers.IndexController{}, "get:Index")
	beego.Router("/message", &controllers.IndexController{}, "get:Message")
	beego.Router("/about", &controllers.IndexController{}, "get:About")
	beego.Router("/user", &controllers.IndexController{}, "get:GetUser")
	beego.Router("/register", &controllers.IndexController{}, "get:GetRegister")
	beego.Router("/reg", &controllers.UserController{}, "post:Register")
	beego.Router("/login", &controllers.UserController{}, "post:Login")
	beego.Router("/logout", &controllers.UserController{}, "get:Logout")
	beego.Router("/details/:key", &controllers.IndexController{}, "get:GetDetail")
	beego.Router("/comment/:key", &controllers.IndexController{}, "get:GetComment")

	beego.Router("/praise/:type/:key", &controllers.PraiseController{}, "get:Praise")

	noteNs := beego.NewNamespace("/note",
		beego.NSRouter("/new", &controllers.NoteController{}, "get:Index"),
		beego.NSRouter("/save/:key", &controllers.NoteController{}, "post:Save"),
		beego.NSRouter("/edit/:key", &controllers.NoteController{}, "get:EditPage"),
		beego.NSRouter("/del/:key", &controllers.NoteController{}, "post:Delete"),
	)

	messageNs := beego.NewNamespace("/message",
		beego.NSRouter("/new/?:key", &controllers.MessageController{}, "post:Save"),
		beego.NSRouter("/count", &controllers.MessageController{}, "get:GetSysMessageCount"),
		beego.NSRouter("/query", &controllers.MessageController{}, "get:GetSysMessages"),
	)
	beego.AddNamespace(noteNs, messageNs)
}
