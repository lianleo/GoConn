package ctrl

import (
	"github.com/gin-gonic/gin"
	"github.com/globalsign/mgo/bson"
	"github.com/lianleo/GoCommon/httpd"
	"github.com/lianleo/GoConn/mongo/service"
)

func AddCollection(c *gin.Context) {
	var req struct {
		Coll string `json:"coll"`
		httpd.RequestValid
	}

	if err := httpd.BindJSON(c, &req); err != nil {
		httpd.WriteParamError(c, err.Error())
		return
	}

	err := service.AddCollection(c, "default", req.Coll)
	if err != nil {
		httpd.WriteError2(c, err)
	} else {
		httpd.WriteOk(c, "OK")
	}
}

func Insert(c *gin.Context) {
	var req struct {
		Coll string `json:"coll"`
		Data bson.M `json:"data"`
		httpd.RequestValid
	}

	if err := httpd.BindJSON(c, &req); err != nil {
		httpd.WriteParamError(c, err.Error())
		return
	}

	err := service.Insert(c, "default", req.Coll, req.Data)
	if err != nil {
		httpd.WriteError2(c, err)
	} else {
		httpd.WriteOk(c, "OK")
	}
}

func Query(c *gin.Context) {
	var req struct {
		Coll string `json:"coll"`
		Data bson.M `json:"data"`
		httpd.RequestValid
	}

	if err := httpd.BindJSON(c, &req); err != nil {
		httpd.WriteParamError(c, err.Error())
		return
	}

	list, err := service.Query(c, "default", req.Coll, req.Data)
	if err != nil {
		httpd.WriteError2(c, err)
	} else {
		httpd.WriteOk(c, httpd.ListResp{Data: list, TotalCount: len(list)})
	}
}
