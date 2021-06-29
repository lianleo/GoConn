package conn

import (
	"fmt"
	"time"

	"github.com/go-redis/redis"
	"github.com/lianleo/GoConn/redis/model"
)

var clientMap map[string]*redis.Client

//init 初始化连接Map
func init() {
	clientMap = make(map[string]*redis.Client)
}

//Connect 连接Redis，创建Redis客户端
func Connect(name string, cfg model.ConnConfig) error {
	client := redis.NewClient(&redis.Options{
		Addr:         cfg.Addr,
		ReadTimeout:  time.Duration(5) * time.Second,
		WriteTimeout: time.Duration(5) * time.Second,
		Password:     cfg.Password,
		DB:           cfg.DB,
	})
	_, err := client.Ping().Result()
	if err != nil {
		return err
	}

	clientMap[name] = client
	return err
}

//GetConnect 获取Redis连接客户端
func GetConnect(name string) (*redis.Client, error) {
	client, ok := clientMap[name]
	if ok {
		return client, nil
	}
	return nil, fmt.Errorf("%s连接不存在", name)
}
