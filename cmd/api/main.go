package main

import (
	"fmt"
	"graze/config"
	"graze/service"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	config.Load()

	e := service.Engine()
	service.Handler(e)
	s := service.New(e, fmt.Sprintf(":%s", config.Conf.Port))
	go signalProcess(s)
	service.Run(s)
}

func signalProcess(srv *http.Server) {
	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGHUP, syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT)

	s := <-c
	switch s {
	case syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT:
		log.Printf("signal is %s", s)
		srv.Close()
		return
	case syscall.SIGHUP:
	default:
		return
	}
}
