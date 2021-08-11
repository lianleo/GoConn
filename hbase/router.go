package hbase

import (
	"github.com/gin-gonic/gin"
	"github.com/lianleo/GoConn/hbase/ctrl"
)

func Install(root *gin.RouterGroup) {
	r := root.Group("hbase")

	r.POST("conn", ctrl.Connect)
	// r.POST("insert", ctrl.Insert)
	// r.POST("query", ctrl.Query)

}
