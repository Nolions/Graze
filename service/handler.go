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

	_, respErr := Client.NewIncident(i.Title, i.Describe, i.Deadline)
	if respErr != nil {
		c.JSON(http.StatusInternalServerError, respErr)
		return
	}
	c.JSON(http.StatusNoContent, nil)
}

// 刪除事件
func DeleteHandler(c *gin.Context) {
	_, resErr := Client.DeleteIncident(c.Param("uid"))
	if resErr != nil {
		resp := errors.ValidatorError{
			Errors: errors.FieldValidatorError(err, i.FieldTrans()),
		}
		resp.Error()
		c.JSON(http.StatusInternalServerError, resp)
		return
	}

	c.JSON(http.StatusNoContent, nil)
}

// 批次刪除
func MultiDeleteHandler(c *gin.Context) {
	uids, ok := c.Request.URL.Query()["uid"]
	if !ok || len(uids) <= 0 {
		resp := errors.ValidatorError{
			Errors:errors.FieldErrorMsg{
				"uid":"至少填入一組事件Id",
			},
		}
		resp.Error()
		c.JSON(http.StatusInternalServerError, resp)
		return
	}

	Client.MultiDeleteIncident(uids)
	c.JSON(http.StatusNoContent, nil)
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
		c.JSON(http.StatusInternalServerError, resp)
		return
	}

	_, respErr := Client.EditIncident(c.Param("uid"), i.Title, i.Describe, i.Deadline)
	if respErr != nil {
		c.JSON(http.StatusInternalServerError, respErr)
		return
	}
	c.JSON(http.StatusNoContent, nil)
}
