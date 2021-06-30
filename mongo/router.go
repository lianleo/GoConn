package mongo

import (
	"github.com/gin-gonic/gin"
	"github.com/lianleo/GoConn/mongo/ctrl"
)

func Install(root *gin.RouterGroup) {
	r := root.Group("mongo")

	r.POST("add/coll", ctrl.AddCollection)
	r.POST("insert", ctrl.Insert)
	r.POST("query", ctrl.Query)

}
