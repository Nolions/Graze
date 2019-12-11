package service

import (
	"github.com/gin-gonic/gin"
	"github.com/gofrs/uuid"
	"log"
	"net/http"
)

func incidentUid(c *gin.Context) {
	uid := c.Param("uid")

	if uid == "" {
		c.AbortWithStatusJSON(http.StatusOK, gin.H{
			"code": 101,
			"msg":  "lack incident's id",
		})
		return
	}

	if len(uid) != 36 {
		c.AbortWithStatusJSON(http.StatusOK, gin.H{
			"code": 102,
			"msg":  "Incident Id Format Error.",
		})
		return
	}

	_, err := uuid.FromString(uid)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusOK, gin.H{
			"code": 101,
			"msg":  "Error Incident Id",
		})
		return
	}
	log.Println("ass")
	c.Next()
}

func incidentParams(c *gin.Context) {
	e := new(Event)
	c.BindJSON(&e)

	var errs []string
	if e.Title == "" {
		errs = append(errs, "title 為必填之項目")
	}

	if e.Describe == "" {
		errs = append(errs, "describe 為必填之項目")
	}

	if len(errs) > 0 {
		c.AbortWithStatusJSON(http.StatusOK, gin.H{
			"code": 101,
			"msg":  "Lack Required Params.",
			"err":  errs,
		})

		return
	}

	c.Next()
	return

}
