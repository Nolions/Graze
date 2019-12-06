package api

import (
	"github.com/gin-gonic/gin"
	"graze/config"
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
	c.String(http.StatusOK, "aa")
	c.String(http.StatusOK, config.Conf.DatastoreHost)
	c.JSON(http.StatusOK, convertList())
}

func CreatorHandler(c *gin.Context) {
	e := new(Event)
	c.BindJSON(&e)

	event := models.New()
	event.Title = e.Title
	event.Describe = e.Describe
	event.Deadline = e.Deadline
	if !event.Store() {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    1005001,
			"message": "Store Error.",
		})
		return
	}


	c.JSON(http.StatusOK, convertList())
}

func DeleteHandler(c *gin.Context) {
	uid := c.Param("uid")
	if _, ok := events[uid]; !ok {
		noDataFound(c)
	}

	delete(events, uid)

	c.JSON(http.StatusOK, convertList())
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

	c.JSON(http.StatusOK, convertList())
}

func noDataFound(c *gin.Context) {
	c.JSON(http.StatusNotFound, gin.H{
		"code":    1004041,
		"message": "No Data Found.",
	})
	return
}

func convertList() []Event {
	if len(events) <= 0 {
		return make([]Event, 0)
	}

	var list []Event
	for _, a := range events {
		list = append(list, a)
	}
	return list
}
