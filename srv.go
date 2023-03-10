package main

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type Srv struct {
	conf    *Config
	httpSrv *http.Server
}

func NewSvr(conf *Config) *Srv {
	srv := &Srv{
		conf: conf,
	}

	gin.SetMode(gin.ReleaseMode)
	// gin.SetMode(gin.DebugMode)

	router := gin.New()
	router.Use(gin.LoggerWithFormatter(func(param gin.LogFormatterParams) string {
		return fmt.Sprintf("[%s] %s %s %d %s %s\n",
			param.TimeStamp.Format("2006-01-02 15:04:05"),
			param.ClientIP,
			param.Method,
			param.StatusCode,
			param.Path,
			param.ErrorMessage,
		)
	}))
	router.Use(gin.Recovery())
	router.StaticFS("/", http.Dir(conf.Path))

	srv.httpSrv = &http.Server{
		Addr:         srv.GetAddr(),
		Handler:      router,
		ReadTimeout:  time.Minute * 10,
		WriteTimeout: time.Minute * 10,
	}
	return srv
}

func (s *Srv) GetAddr() string {
	return fmt.Sprintf("%s:%d", s.conf.Ip, s.conf.Port)
}

func (s *Srv) Start() {
	go func() {
		err := s.httpSrv.ListenAndServe()
		if err != nil && err != http.ErrServerClosed {
			panic(err)
		}
	}()
}

func (s *Srv) Shutdown(ctx context.Context) error {
	return s.httpSrv.Shutdown(ctx)
}
