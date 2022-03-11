package controllers

import (
	"my-blog/models"
	"my-blog/syserror"
)

type MessageController struct {
	BaseController
}

func (this *MessageController) Save() {
	this.MustLogin()
	key := this.Ctx.Input.Param(":key")
	content := this.GetMustString("content", "评论内容不能为空")
	//_, err := models.QueryNoteByKey(key)
	//if err != nil {
	//	this.Abort500(syserror.New("文章不存在", err))
	//}
	messageKey := this.UUID()
	message := models.Message{
		MessageKey: messageKey,
		Content:    content,
		User:       this.User,
		UserId:     this.User.ID,
		NoteKey:    key,
	}
	err := models.SaveMessage(&message)
	if err != nil {
		this.Abort500(syserror.New("评论失败", err))
	}
	if len(key) == 0 {
		this.JsonOkH("留言成功", H{"data": message})
	} else {
		this.JsonOk("评论成功", "/details/"+key)
	}
}

func (this *MessageController) GetSysMessageCount() {
	count, err := models.QueryMessageCountByNoteKey("")
	if err != nil {
		this.Abort500(syserror.New("查询失败", err))
	}
	this.JsonOkH("查询成功", H{"count": count})
}

func (this *MessageController) GetSysMessages() {
	page, err := this.GetInt("pageno", 1)
	if err != nil || page <= 0 {
		page = 1
	}
	pageSize, err := this.GetInt("pagesize", 10)
	if err != nil || pageSize <= 0 {
		pageSize = 3
	}
	ms, err := models.QueryPageMessageByNoteKey("", page, pageSize)
	if err != nil {
		this.Abort500(syserror.New("查询失败", err))
	}
	this.JsonOkH("查询成功", H{"data": ms})
}
