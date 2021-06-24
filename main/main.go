package main

import (
	"fmt"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/lianleo/GoCommon/log"
	"github.com/lianleo/GoConn/config"
	"github.com/lianleo/GoConn/router"
)

func init() {
	configPath := ""
	if len(os.Args) > 1 {
		configPath = os.Args[1]
	}

	if configPath == "" {
		configPath = os.Getenv("WEB_APP_CONFIG")
	}

	fmt.Println(configPath)
	err := config.Install(configPath)
	if err != nil {
		log.Error("init error", err)
		panic(err)
	}
}

func main() {
	log.Infof("start %s", config.Params.WebAPP.App)
	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Wellcome to Go Conn",
			"time":    time.Now().Format("2006-01-02 15:04:05.000"),
		})
	})

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
			"time":    time.Now().Format("2006-01-02 15:04:05.000"),
		})
	})

	router.Install(r)

	log.Infof("%s service begin run", config.Params.WebAPP.App)

	r.Run(":" + config.Params.WebAPP.Port)
}
