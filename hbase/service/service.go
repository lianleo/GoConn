package service

import (
	"context"

	"github.com/gin-gonic/gin"
	"github.com/lianleo/GoConn/global"
	"github.com/lianleo/GoConn/hbase/conn"
	"github.com/lianleo/GoConn/hbase/model"
)

const (
	HBaseCookieKey = "CurrentHBaseConnection"
)

func Connect(ctx context.Context, name string, config model.HBaseConfig) error {
	err := conn.Connect(name, config)
	if err != nil {
		return err
	}

	ctx.(*gin.Context).SetCookie(HBaseCookieKey, name, global.Config.WebAPP.Expires, "", global.Config.WebAPP.Domain, false, true)
	return nil
}
