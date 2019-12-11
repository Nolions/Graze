package api

import (
	"github.com/gin-gonic/gin"
	"graze/models"
	"net/http"
	"time"
)

type Event struct {
	models.Incident
	Uid      string    `json:"uid"`
	CreateAt time.Time `json:"create_at"`
}

var Client *models.Datastore

// 所有事件
func ListHandler(c *gin.Context) {
	list := Client.AllIncident()

	if len(list) > 0 {
		c.JSON(http.StatusOK, list)
	} else {
		c.JSON(http.StatusOK, gin.H{})
	}
}

// 新增事件
func CreatorHandler(c *gin.Context) {
	e := new(Event)
	c.BindJSON(&e)

	i := new(models.Incident)
	i.New()
	i.Title = e.Title
	i.Describe = e.Describe
	i.Deadline = e.Deadline

	Client.NewIncident(i)

	c.JSON(http.StatusNoContent, nil)
}

// 刪除事件
func DeleteHandler(c *gin.Context) {
	Client.DeleteIncident(c.Param("uid"))

	c.JSON(http.StatusNoContent, nil)
}

// 編輯事件
func EditHandler(c *gin.Context) {
	e := new(Event)
	c.BindJSON(&e)

	Client.EditIncident(c.Param("uid"), e.Title, e.Describe, e.Deadline)
}