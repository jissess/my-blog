package controllers

import (
	"my-blog/models"
	"my-blog/syserror"
)

type IndexController struct {
	BaseController
}

func (this *IndexController) Index() {
	limit := 3
	page, err := this.GetInt("page", 1)
	if err != nil || page <= 0 {
		page = 1
	}
	title := this.GetString("title")
	notes, err := models.QueryNotesByPage(title, page, limit)
	count, err := models.QueryNotesCount(title)
	if err != nil {
		this.Abort500(err)
	}
	pages := count / limit
	if count%limit != 0 {
		pages += 1
	}
	this.Data["pages"] = pages
	this.Data["notes"] = notes
	this.Data["page"] = page
	this.Data["title"] = title
	this.TplName = "index.html"
}

func (this *IndexController) Message() {
	this.TplName = "message.html"
}

func (this *IndexController) About() {
	this.TplName = "about.html"
}

func (this *IndexController) GetUser() {
	this.TplName = "user.html"
}

func (this *IndexController) GetRegister() {
	this.TplName = "register.html"
}

func (this *IndexController) GetDetail() {
	key := this.Ctx.Input.Param(":key")
	note, err := models.QueryNoteByKey(key)
	if err != nil {
		this.Abort500(syserror.New("文章不存在", err))
	}
	models.NoteVisitInc(key)
	messages, err := models.QueryMessageByNoteKey(key)
	if err != nil {
		this.Abort500(syserror.New("暂无评论", err))
	}
	this.Data["note"] = note
	this.Data["messages"] = messages
	this.TplName = "detail.html"
}

func (this *IndexController) GetComment() {
	key := this.Ctx.Input.Param(":key")
	note, err := models.QueryNoteByKey(key)
	if err != nil {
		this.Abort500(syserror.New("文章不存在", err))
	}
	this.Data["note"] = note
	this.TplName = "comment.html"
}
