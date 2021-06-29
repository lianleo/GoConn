package redis

import (
	"github.com/gin-gonic/gin"
	"github.com/lianleo/GoConn/redis/ctrl"
)

func Install(root *gin.RouterGroup) {
	r := root.Group("redis")

	r.POST("connect", ctrl.Connect)
	r.GET("get/:key", ctrl.GetKey)
	r.POST("set", ctrl.SetKey)

}
