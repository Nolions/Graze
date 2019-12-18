package service

import (
	"github.com/gin-gonic/gin"
	"graze/errors"
	"graze/models"
	"net/http"
)

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
	i := new(models.Incident)
	c.BindJSON(&i)

	err := errors.Validate.Struct(i)
	if err != nil {
		resp := errors.ValidatorError{
			Errors: errors.FieldValidatorError(err, i.FieldTrans()),
		}
		resp.Error()
		c.JSON(500, resp)
		return
	}

	Client.NewIncident(i.Title, i.Describe, i.Deadline)
	c.JSON(http.StatusNoContent, nil)
}

// 刪除事件
func DeleteHandler(c *gin.Context) {
	Client.DeleteIncident(c.Param("uid"))
	c.JSON(http.StatusNoContent, nil)
}

func MultiDeleteHandler(c *gin.Context)  {
	
}

// 編輯事件
func EditHandler(c *gin.Context) {
	i := new(models.Incident)
	c.BindJSON(&i)

	err := errors.Validate.Struct(i)
	if err != nil {
		resp := errors.ValidatorError{
			Errors: errors.FieldValidatorError(err, i.FieldTrans()),
		}
		resp.Error()
		c.JSON(500, resp)
		return
	}

	Client.EditIncident(c.Param("uid"), i.Title, i.Describe, i.Deadline)
	c.JSON(http.StatusNoContent, nil)
}
