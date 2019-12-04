package main

import (
	"fmt"
	"graze/config"
	"graze/server"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	config.Load()

	e := server.Engine()

	server.Handler(e)
	s := server.New(e, fmt.Sprintf(":%s", config.Conf.Port))
	go signalProcess(s)
	server.Run(s)
}

func signalProcess(srv *http.Server) {
	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGHUP, syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT)

	s:= <-c
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
