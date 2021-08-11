package ctrl

import (
	"github.com/gin-gonic/gin"
	"github.com/lianleo/GoCommon/httpd"
	"github.com/lianleo/GoConn/hbase/model"
	"github.com/lianleo/GoConn/hbase/service"
)

func Connect(c *gin.Context) {
	var req struct {
		Name   string            `json:"name"`
		Config model.HBaseConfig `json:"config"`
		httpd.RequestValid
	}
	if err := httpd.BindJSON(c, &req); err != nil {
		httpd.WriteParamError(c, err.Error())
		return
	}
	if err := service.Connect(c, req.Name, req.Config); err != nil {
		httpd.WriteParamError(c, err.Error())
		return
	}
	httpd.WriteOk(c, "ok")
}
