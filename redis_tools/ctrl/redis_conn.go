package ctrl

import (
	"github.com/gin-gonic/gin"
	"github.com/lianleo/GoCommon/httpd"
	"github.com/lianleo/GoCommon/log"
	"github.com/lianleo/GoCommon/tools"
)

func Redis(c *gin.Context) {
	log.Info("hello ")
	rs := struct {
		Message string `json:"message"`
		Time    string `json:"time"`
	}{
		"Hello",
		tools.CNDateTimeFormat(tools.NowEpochMS(), "2006-01-02 15:04:05"),
	}
	httpd.WriteOk(c, rs)
}
