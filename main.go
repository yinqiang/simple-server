package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/gin-gonic/gin"
)

var (
	conf *Config
)

func main() {
	var err error
	conf, err = loadConf("./conf.json")
	if err != nil {
		panic(err)
	}

	gin.SetMode(gin.ReleaseMode)

	router := gin.New()
	router.Static("/", conf.StaticPath)

	srv := &http.Server{
		Addr:           fmt.Sprintf("localhost:%d", 9090),
		Handler:        router,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	go func() {
		err := srv.ListenAndServe()
		if err != nil && err != http.ErrServerClosed {
			panic(err)
		}
	}()
	log.Printf("server on: http://%s:%d/\n", conf.Ip, conf.Port)

	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	<-quit
	log.Println("server shutdown")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err = srv.Shutdown(ctx); err != nil {
		log.Fatal("server shutdown, error:", err)
	}
	log.Println("server exit")
}
