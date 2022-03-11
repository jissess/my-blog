package controllers

import (
	"errors"
	"my-blog/models"
	"my-blog/syserror"
	"strings"
)

type UserController struct {
	BaseController
}

func (this *UserController) Register() {
	name := this.GetMustString("name", "昵称不能为空")
	email := this.GetMustString("email", "邮箱不能为空")
	password := this.GetMustString("password", "密码不能为空")
	rePassword := this.GetMustString("re-password", "密码不能为空")
	if strings.Compare(password, rePassword) != 0 {
		this.Abort500(errors.New("两次输入密码不一致"))
	}
	if u, err := models.QueryUserByName(name); err == nil && u.ID > 0 {
		this.Abort500(errors.New("用户昵称已存在"))
	}
	if u, err := models.QueryUserByEmail(email); err == nil && u.ID > 0 {
		this.Abort500(errors.New("用户邮箱已注册"))
	}
	var user models.User
	user.Name = name
	user.Password = password
	user.Email = email
	user.Avatar = "/static/images/info-img.png"
	if err := models.SaveUser(&user); err != nil {
		this.Abort500(syserror.New("用户保存失败", err))
	}
	this.JsonOk("注册成功", "/user")
}

func (this *UserController) Login() {
	email := this.GetMustString("email", "邮箱不能为空")
	password := this.GetMustString("password", "密码不能为空")
	user, err := models.QueryUserByEmailAndPwd(email, password)
	if err != nil {
		this.Abort500(syserror.New("登录失败", err))
	}
	this.SetSession(SESSION_USER_KEY, user)
	this.JsonOk("登录成功", "/")
}

func (this *UserController) Logout() {
	this.MustLogin()
	this.DelSession(SESSION_USER_KEY)
	this.Redirect("/", 302)
}
