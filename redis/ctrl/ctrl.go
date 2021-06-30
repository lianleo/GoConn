package ctrl

import (
	"github.com/gin-gonic/gin"
	"github.com/lianleo/GoCommon/httpd"
	"github.com/lianleo/GoConn/redis/model"
	"github.com/lianleo/GoConn/redis/service"
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

func GetKey(c *gin.Context) {
	key := c.Query("key")
	if len(key) == 0 {
		key = c.Params.ByName("key")
	}

	connName, err := c.Cookie(service.RedisCookieKey)
	if err != nil {
		httpd.WriteError2(c, err)
	}

	result, err := service.GetKey(c, connName, key)
	if err != nil {
		httpd.WriteError2(c, err)
	} else {
		httpd.WriteOk(c, result)
	}

}

func SetKey(c *gin.Context) {
	var req struct {
		Key   string `json:"key"`
		Value string `json:"value"`
		httpd.RequestValid
	}

	if err := httpd.BindJSON(c, &req); err != nil {
		httpd.WriteParamError(c, err.Error())
		return
	}

	connName, err := c.Cookie(service.RedisCookieKey)
	if err != nil {
		httpd.WriteError2(c, err)
	}

	result, err := service.SetKey(c, connName, req.Key, req.Value)
	if err != nil {
		httpd.WriteError2(c, err)
	} else {
		httpd.WriteOk(c, result)
	}
}

func Keys(c *gin.Context) {
	pattern := c.Query("pattern")
	if len(pattern) == 0 {
		pattern = c.Params.ByName("pattern")
	}

	connName, err := c.Cookie(service.RedisCookieKey)
	if err != nil {
		httpd.WriteError2(c, err)
	}

	result, err := service.Keys(c, connName, pattern)
	if err != nil {
		httpd.WriteError2(c, err)
	} else {
		httpd.WriteOk(c, result)
	}
}
