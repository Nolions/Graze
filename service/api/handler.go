package api

import (
	"github.com/gin-gonic/gin"
	"graze/models"
	"net/http"
	"time"
)

type Event struct {
	models.Event
	Uid      string    `json:"uid"`
	CreateAt time.Time `json:"create_at"`
}

var events = make(map[string]Event)

// 所有事件
func ListHandler(c *gin.Context) {
	list := new(Event).All()
	if len(list) > 0 {
		c.JSON(http.StatusOK, new(Event).All())
	} else {
		c.JSON(http.StatusOK, gin.H{})
	}
}

// 新增事件
func CreatorHandler(c *gin.Context) {
	e := new(Event)
	c.BindJSON(&e)

	event := models.New()
	event.Title = e.Title
	event.Describe = e.Describe
	event.Deadline = e.Deadline
	if !event.Creator() {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    1005001,
			"message": "Store Error.",
		})
		return
	}

	c.JSON(http.StatusNoContent, nil)
}

// 刪除事件
func DeleteHandler(c *gin.Context) {
	uid := c.Param("uid")

	var e = new(models.Event)
	e.Uid = uid
	e.Delete()

	c.JSON(http.StatusNoContent, nil)
}

// 編輯事件
func EditHandler(c *gin.Context) {
	e := new(Event)
	c.BindJSON(&e)

	var event = new(models.Event)
	event.Uid = c.Param("uid")
	event.Title = e.Title
	event.Describe = e.Describe
	event.Deadline = e.Deadline
	event.Edit()
	c.JSON(http.StatusNoContent, nil)
}