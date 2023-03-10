package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"os"
	"os/signal"
	"time"
)

const (
	DefaultHost = "localhost"
	DefaultPort = 9980
	DefaultPath = "./"
)

func main() {
	defer func() {
		err := recover()
		if err != nil {
			log.Printf("Error: %v\n", err)
		}
	}()

	ip := flag.String("ip", "", fmt.Sprintf("ip address, default: %s", DefaultHost))
	port := flag.Int("port", 0, fmt.Sprintf("port, default: %d", DefaultPort))
	rootPath := flag.String("root", "", fmt.Sprintf("root path, default: %s", DefaultPath))
	flag.Parse()

	confPath := "./svr.conf"
	conf, err := LoadConfigOrDefault(confPath, DefaultPath, DefaultHost, DefaultPort)
	if err != nil {
		panic(err)
	}

	if len(*ip) > 0 {
		conf.Ip = *ip
	}
	if *port > 0 {
		conf.Port = *port
	}
	if len(*rootPath) > 0 {
		conf.Path = *rootPath
	}
	conf.Save(confPath)

	srv := NewSvr(conf)
	srv.Start()
	log.Printf("dir: %s", conf.Path)
	log.Printf("server on: http://%s/\n", srv.GetAddr())

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit
	log.Println("server shutdown")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("server shutdown, error:", err)
	}
	log.Println("server exit")
}
