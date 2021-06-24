package router

import (
	"github.com/gin-gonic/gin"
	"github.com/lianleo/GoCommon/httpd"
	"github.com/lianleo/GoCommon/log"
	"github.com/lianleo/GoCommon/tools"
	"github.com/lianleo/GoConn/config"
	redis_tools "github.com/lianleo/GoConn/redis_tools/ctrl"
)

//Install 初始化路由
func Install(eng *gin.Engine) {
	gocon := eng.Group(config.Params.WebAPP.App)
	gocon.Use(middleware())

	gocon.GET("/ping", func(c *gin.Context) {
		for i := 0; i < int(1000); i++ {
			log.Info("测试日志")
		}
		c.JSON(200, gin.H{
			"project": "go connection ",
			"time":    tools.NowEpochMS(),
			"message": "pong",
		})
	})

	v1 := gocon.Group("api/v1")

	v1.Static("/static", config.Params.WebAPP.StaticDir)

	v1.GET("redis", redis_tools.Redis)

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
