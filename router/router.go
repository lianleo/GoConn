package router

import (
	"github.com/gin-gonic/gin"
	"github.com/lianleo/GoCommon/httpd"
	"github.com/lianleo/GoCommon/log"
	"github.com/lianleo/GoCommon/tools"
	"github.com/lianleo/GoConn/global"
	"github.com/lianleo/GoConn/mysql"
	"github.com/lianleo/GoConn/redis"
)

//Install 初始化路由
func Install(eng *gin.Engine) {
	goconn := eng.Group(global.Config.WebAPP.App)
	goconn.Use(middleware())

	goconn.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"project": "go connection ",
			"time":    tools.NowEpochMS(),
			"message": "pong",
		})
	})

	v1 := goconn.Group("api/v1")

	v1.Static("/static", global.Config.WebAPP.StaticDir)

	redis.Install(v1)

	mysql.Install(v1)

}

func middleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()
		writeRecord(c) //记录访问日志
	}
}

func writeRecord(c *gin.Context) {
	bc := httpd.BC(c)
	v, _ := bc.GetBindJSON()

	record := struct {
		Method        string      `bson:"method" json:"method"`
		RequestIP     string      `bson:"req_ip" json:"req_id"`
		URL           string      `bson:"url" json:"url"`
		RequestHeader interface{} `bson:"req_header" json:"req_header"`
		RequestCookie interface{} `bson:"req_cookie" json:"req_cookie"`
		RequestQuery  interface{} `bson:"req_query" json:"req_query"`
		RequestParam  interface{} `bson:"req_param" json:"req_param"`
		RequestBody   interface{} `bson:"req_body" json:"req_body"`
		ResponseCode  interface{} `bson:"resp_code" json:"resp_code"`
		ResponseData  interface{} `bson:"resp_data" json:"resp_data"`
	}{
		Method:        c.Request.Method,
		RequestIP:     bc.RemoteAddr(),
		URL:           c.Request.URL.Path,
		RequestHeader: c.Request.Header,
		RequestCookie: c.Request.Cookies(),
		RequestQuery:  c.Request.URL.RawQuery,
		RequestParam:  c.Params,
		RequestBody:   v,
		// ResponseCode:  respCode,
		// ResponseData:  resp,
	}

	log.Info(record)
}
