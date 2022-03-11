package controllers

import (
	"errors"
	"my-blog/models"
	"my-blog/syserror"
)

type PraiseController struct {
	BaseController
}

func (this *PraiseController) NextPrepare() {
	this.MustLogin()
}

func (this *PraiseController) Praise() {
	pType := this.Ctx.Input.Param(":type")
	key := this.Ctx.Input.Param(":key")
	table := "notes"
	columnKey := ""
	switch pType {
	case "note":
		table = "notes"
		columnKey = "note_key"
	case "message":
		table = "messages"
		columnKey = "message_key"
	default:
		this.Abort500(errors.New("未知类型"))
	}
	praise, err := models.UpdatePraise(table, columnKey, key, this.User.ID)
	if err != nil {
		if err1, ok := err.(syserror.HasPraiseError); ok {
			this.Abort500(err1)
		}
		this.Abort500(syserror.New("点赞失败", err))
	}
	this.JsonOkH("点赞成功", H{"praise": praise})
}
