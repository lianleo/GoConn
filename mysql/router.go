package mysql

import (
	"github.com/gin-gonic/gin"
	"github.com/lianleo/GoConn/mysql/ctrl"
)

func Install(root *gin.RouterGroup) {
	r := root.Group("mysql")

	r.POST("connect", ctrl.Connect)
	r.POST("insert", ctrl.Insert)
	r.POST("run/sql", ctrl.RunSQL)

}
