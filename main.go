package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/op-y/dnspod/application/controller"
	"github.com/op-y/dnspod/config"

	"github.com/gin-gonic/gin"
)

func main() {
	log.Printf("===System Startup===")
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	// gin
	routes := gin.Default()
	log.Printf("gin will start with port:%s", config.CFG.Port)
	go controller.StartGin(config.CFG.Port, routes)

	if config.CFG.LogLevel != "debug" {
		gin.SetMode(gin.ReleaseMode)
	}
	gin.DisableConsoleColor()
	//f, _ := os.Create("gin.log")
	//gin.DefaultWriter = io.MultiWriter(f)
	//gin.DefaultWriter = io.MultiWriter(f, os.Stdout)

	// waiting
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)
	go func() {
		<-sigs
		fmt.Println()
		os.Exit(0)
	}()
	select {}
}
