package controllers

import (
	"github.com/astaxie/beego/logs"
	"my-blog/syserror"
)

type ErrorController struct {
	BaseController
}

func (c *ErrorController) Error404() {
	c.TplName = "error/404.html"
	if c.IsAjax() {
		c.jsonError(syserror.Error404{})
	}
}

func (c *ErrorController) Error500() {
	c.TplName = "error/500.html"
	err, ok := c.Data["error"].(error)
	if !ok {
		err = syserror.New("未知错误", nil)
	}
	serr, ok := err.(syserror.Error)
	if !ok {
		serr = syserror.New(err.Error(), nil)
	}
	if serr.ReasonError() != nil {
		logs.Info(serr.Error(), serr.ReasonError())
	}
	if c.IsAjax() {
		c.jsonError(serr)
	} else {
		c.Data["content"] = serr.Error()
	}
}

func (c *ErrorController) jsonError(err syserror.Error) {
	c.Ctx.Output.Status = 200
	c.Data["json"] = map[string]interface{}{"code": err.Code(), "msg": err.Error()}
	c.ServeJSON()
}
