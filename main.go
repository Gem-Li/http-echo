// Author: f210 (508699929@qq.com)
// Date: 2025/09/02
// Description main.go:

package main

import (
	"flag"
	"fmt"
	"http-echo/lib/buildinfo"
	"log"
	"time"

	"github.com/gin-gonic/gin"
)

var (
	port int
	help bool
)

func initFlag() {
	flag.IntVar(&port, "p", 8181, "set http listen port")
	flag.BoolVar(&help, "h", false, "show help message")
	flag.Parse()
}

func main() {

	initFlag()

	buildinfo.InitBuildInfo()

	if help {
		flag.Usage()
		return
	}

	ginLogFormat := gin.LogFormatter(func(param gin.LogFormatterParams) string {
		return fmt.Sprintf("%s %s %s %s %s %d %s %s %s\n",
			param.TimeStamp.Format(time.RFC3339),
			param.ClientIP,
			param.Method,
			param.Path,
			param.Request.Proto,
			param.StatusCode,
			param.Latency,
			param.Request.UserAgent(),
			param.ErrorMessage,
		)
	})
	ginLogConfig := gin.LoggerConfig{
		Formatter: ginLogFormat,
		SkipPaths: []string{},
	}

	gin.SetMode(gin.ReleaseMode)
	r := gin.New()

	r.Use(gin.Recovery())
	r.Use(gin.LoggerWithConfig(ginLogConfig))

	RegistRoute(r)

	log.Println("http server start...")
	log.Fatal(r.Run(fmt.Sprintf(":%d", port)))
}
