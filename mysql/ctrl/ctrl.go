package ctrl

import (
	"github.com/gin-gonic/gin"
	"github.com/lianleo/GoCommon/httpd"
	"github.com/lianleo/GoConn/mysql/model"
	"github.com/lianleo/GoConn/mysql/service"
)

func Connect(c *gin.Context) {
	var req struct {
		Name   string           `json:"name"`
		Config model.ConnConfig `json:"config"`
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

func Insert(c *gin.Context) {
	var req struct {
		Key   string `json:"key"`
		Value string `json:"value"`
		httpd.RequestValid
	}

	if err := httpd.BindJSON(c, &req); err != nil {
		httpd.WriteParamError(c, err.Error())
		return
	}

	connName, err := c.Cookie(service.MysqlCookieKey)
	if err != nil {
		httpd.WriteError2(c, err)
	}

	err = service.Insert(c, connName, req)
	if err != nil {
		httpd.WriteError2(c, err)
	} else {
		httpd.WriteOk(c, "ok")
	}

}

func RunSQL(c *gin.Context) {
	var req struct {
		SQL string `json:"sql"`
		httpd.RequestValid
	}

	if err := httpd.BindJSON(c, &req); err != nil {
		httpd.WriteParamError(c, err.Error())
		return
	}

	connName, err := c.Cookie(service.MysqlCookieKey)
	if err != nil {
		httpd.WriteError2(c, err)
	}

	rs, err := service.RunSQL(c, connName, req.SQL)
	if err != nil {
		httpd.WriteError2(c, err)
	} else {
		httpd.WriteOk(c, rs)
	}
}
