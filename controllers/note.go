package controllers

import (
	"bytes"
	"errors"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"github.com/jinzhu/gorm"
	"my-blog/models"
	"my-blog/syserror"
)

type NoteController struct {
	BaseController
}

func (this *NoteController) NextPrepare() {
	this.MustLogin()
	if this.User.Role != 0 {
		this.Abort500(errors.New("权限不足"))
	}
}

func (this *NoteController) Index() {
	this.Data["key"] = this.UUID()
	this.TplName = "note_new.html"
}

func (this *NoteController) EditPage() {
	key := this.Ctx.Input.Param(":key")
	note, err := models.QueryNoteByKeyAndUserId(key, this.User.ID)
	if err != nil {
		this.Abort500(syserror.New("文章不存在", err))
	}
	this.Data["note"] = note
	this.Data["key"] = key
	this.TplName = "note_edit.html"
}

func (this *NoteController) Save() {
	key := this.Ctx.Input.Param(":key")
	title := this.GetMustString("title", "标题不能为空，请输入标题")
	content := this.GetMustString("content", "内容不能为空，请输入文章内容")
	noteInfo, err := models.QueryNoteByKey(key)
	var note models.Note
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			note.NoteKey = key
			note.Title = title
			note.Content = content
			note.UserId = int(this.User.ID)
			note.User = this.User
			if err = models.SaveNote(&note); err != nil {
				this.Abort500(syserror.New("保存失败", err))
			}
			this.JsonOk("保存成功", "/")
		} else {
			this.Abort500(syserror.New("保存失败", err))
		}
	} else {
		noteInfo.Title = title
		noteInfo.Content = content
		note = noteInfo
	}
	note.Summary, _ = getSummary(content)
	if err = models.SaveNote(&note); err != nil {
		this.Abort500(syserror.New("保存失败", err))
	}
	this.JsonOk("保存成功", fmt.Sprintf("/details/%s", key))
}

func (this *NoteController) Delete() {
	key := this.Ctx.Input.Param(":key")
	err := models.DeleteNoteByKeyAndUserId(key, this.User.ID)
	if err != nil {
		this.Abort500(syserror.New("删除失败", err))
	}
	this.JsonOk("删除成功", "/")
}

func getSummary(html string) (string, error) {
	var bufBytes bytes.Buffer
	if _, err := bufBytes.Write([]byte(html)); err != nil {
		return "", err
	}
	doc, err := goquery.NewDocumentFromReader(&bufBytes)
	if err != nil {
		return "", err
	}
	htmlStr := doc.Find("body").Text()
	if len(htmlStr) > 600 {
		htmlStr = htmlStr[:600]
	}
	return htmlStr, nil
}
