package service

import (
	"github.com/gin-gonic/gin"
	"graze/config"
	"graze/models"
	"graze/service/api"
	"log"
	"net/http"
	"strconv"
	"time"
)

func New(r *gin.Engine, addr string) *http.Server {
	log.Printf("Listening on %s", addr)
	return &http.Server{
		Addr:         addr,
		Handler:      r,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}
}

func Handler(router *gin.Engine) {
	api.Client = (new(models.Datastore)).NewClient()

	router.GET("/", api.ListHandler)
	router.POST("/", api.CreatorHandler)
	router.DELETE("/:uid", api.DeleteHandler)
	router.PUT("/:uid", api.EditHandler)
}

func Run(s *http.Server) {
	if err := s.ListenAndServe(); err != nil {
		log.Println(err)
	}
}

func indexHandler(c *gin.Context) {
	c.String(http.StatusOK, strconv.FormatBool(config.Conf.Debug))
}
