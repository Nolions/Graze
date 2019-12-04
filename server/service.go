package server

import (
	"github.com/gin-gonic/gin"
	"graze/config"
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
	router.GET("/", ListHandler)
	router.POST("/", CreatorHandler)
	router.DELETE("/:uid", DeleteHandler)
	router.PUT("/:uid", EditHandler)
}

func Run(s *http.Server) {
	if err := s.ListenAndServe(); err != nil {
		log.Println(err)
	}
}

func indexHandler(c *gin.Context) {
	c.String(http.StatusOK, strconv.FormatBool(config.Conf.Debug))
}
