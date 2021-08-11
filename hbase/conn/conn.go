package conn

import (
	"context"
	"fmt"

	"github.com/lianleo/GoCommon/log"
	"github.com/lianleo/GoConn/hbase/model"
	"github.com/tsuna/gohbase"
	"github.com/tsuna/gohbase/hrpc"
)

var sessionMap map[string]*gohbase.Client

//init 初始化连接Map
func init() {
	sessionMap = make(map[string]*gohbase.Client)
}

func Install() {
	client := gohbase.NewClient("localhost")
	sessionMap["default"] = &client
}

//Connect 连接HBase
func Connect(name string, cfg model.HBaseConfig) error {

	client := gohbase.NewClient(cfg.IP)
	sessionMap[name] = &client

	getRequest, err := hrpc.NewGetStr(context.Background(), "user", "1001")
	if err != nil {
		log.Error(err)
	}
	getRsp, err := client.Get(getRequest)
	if err != nil {
		log.Error(err)
	}

	log.Info(getRsp)

	// admin := gohbase.NewAdminClient("master")
	// columns := []string{"cf"}
	// createReq := hrpc.NewCreateTable(context.Background(), []byte("goTest"), columns)
	// err := admin.CreateTable(createReq)
	// if err != nil {
	// 	log.Error("err:", err)
	// }

	return nil
}

//GetConnect 获取数据库连接客户端
func GetConnect(name string) (*gohbase.Client, error) {
	sess, ok := sessionMap[name]
	if ok {
		return sess, nil
	}
	return nil, fmt.Errorf("%s连接不存在", name)
}
