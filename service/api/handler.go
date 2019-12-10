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

func ListHandler(c *gin.Context) {
	list := new(Event).All()
	if len(list) > 0 {
		c.JSON(http.StatusOK, new(Event).All())
	} else {
		c.JSON(http.StatusOK, gin.H{})
	}
}

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

func DeleteHandler(c *gin.Context) {
	uid := c.Param("uid")

	var e = new(models.Event)
	e.Uid = uid
	e.Delete()

	c.JSON(http.StatusNoContent, nil)
}

func EditHandler(c *gin.Context) {
	uid := c.Param("uid")
	if _, ok := events[uid]; !ok {
		noDataFound(c)
	}

	re := new(Event)
	c.BindJSON(&re)
	e := events[uid]
	e.Title = re.Title
	e.Describe = re.Describe
	//e.Deadline = re.Deadline
	events[uid] = e

	c.JSON(http.StatusNoContent, nil)
}

func noDataFound(c *gin.Context) {
	c.JSON(http.StatusNotFound, gin.H{
		"code":    1004041,
		"message": "No Data Found.",
	})
	return
}