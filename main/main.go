package main

import (
	"fmt"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/lianleo/GoCommon/log"
	"github.com/lianleo/GoConn/global"
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
	err := global.Install(configPath)
	if err != nil {
		log.Error("init error", err)
		panic(err)
	}
}

func main() {
	log.Infof("start %s", global.Config.WebAPP.Title)
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

	log.Infof("%s service begin run", global.Config.WebAPP.Title)

	r.Run(":" + global.Config.WebAPP.Port)
}
