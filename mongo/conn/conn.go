package conn

import (
	"fmt"
	"time"

	"github.com/globalsign/mgo"
	"github.com/lianleo/GoCommon/db/mongo"
)

var sessionMap map[string]*mgo.Database

//init 初始化连接Map
func init() {
	sessionMap = make(map[string]*mgo.Database)
}

func Install(cfg mongo.Config) error {
	return Connect("default", cfg)
}

//Connect 连接MongoDB，创建MongoDB客户端
func Connect(name string, cfg mongo.Config) error {
	var url string
	if cfg.UserName != "" {
		url = fmt.Sprintf("mongodb://%s:%s@%s:%d", cfg.UserName, cfg.Password, cfg.IP, cfg.Port)
	} else {
		url = "mongodb://"
	}
	url += fmt.Sprintf("/%s?%s", cfg.Database, "authSource=admin&readPreference=primary")

	timeout := time.Duration(cfg.Timeout) * time.Second
	// fmt.Println(url)
	session, err := mgo.DialWithTimeout(url, timeout)
	session.SetPoolLimit(cfg.PoolLimit)

	if err != nil {
		return err
	}

	sessionMap[name] = session.DB(cfg.Database)

	return nil
}

//GetConnect 获取数据库连接客户端
func GetConnect(name string) (*mgo.Database, error) {
	sess, ok := sessionMap[name]
	if ok {
		return sess, nil
	}
	return nil, fmt.Errorf("%s连接不存在", name)
}
