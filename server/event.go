package server

import (
	"github.com/gin-gonic/gin"
	"github.com/gofrs/uuid"
	"net/http"
	"time"
)

type Event struct {
	Title    string    `json:"title"`
	Describe string    `json:"describe"`
	Deadline time.Time `json:"deadline"`
}

//
type ResponseEvent struct {
	Event
	Uid      string    `json:"uid"`
	CreateAt time.Time `json:"create_at"`
}

var events = make(map[string]ResponseEvent)

func new() Event {
	return Event{
		Title:    "",
		Describe: "",
		Deadline: time.Time{},
	}
}

func ListHandler(c *gin.Context) {
	c.JSON(http.StatusOK, convertList())
}

func CreatorHandler(c *gin.Context) {
	e := new()
	c.BindJSON(&e)

	uid := uuid.Must(uuid.NewV4()).String()
	events[uid] = ResponseEvent{
		Event: Event{
			Title:    e.Title,
			Describe: e.Describe,
			Deadline: e.Deadline,
		},
		Uid:      uid,
		CreateAt: time.Now(),
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

	re := new()
	c.BindJSON(&re)
	e := events[uid]
	e.Title = re.Title
	e.Describe = re.Describe
	e.Deadline = re.Deadline
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

func convertList() []ResponseEvent {
	if len(events) <= 0 {
		return make([]ResponseEvent, 0)
	}

	var list []ResponseEvent
	for _, a := range events {
		list = append(list, a)
	}
	return list
}
