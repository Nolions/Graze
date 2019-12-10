package service

import (
	"github.com/gin-gonic/gin"
	"graze/config"
)

func Engine() *gin.Engine {
	e := &gin.Engine{}
	if !config.Conf.Debug {
		gin.SetMode(gin.DebugMode)
		e = gin.Default()
	} else {
		gin.SetMode(gin.ReleaseMode)
		e = gin.New()
	}

	return e
}
