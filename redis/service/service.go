package service

import (
	"context"

	"github.com/gin-gonic/gin"
	"github.com/lianleo/GoConn/global"
	"github.com/lianleo/GoConn/redis/conn"
	"github.com/lianleo/GoConn/redis/model"
)

const (
	RedisCookieKey = "CurrentRedisConnection"
)

//Connect 建立Redis连接
func Connect(ctx context.Context, name string, config model.ConnConfig) error {
	err := conn.Connect(name, config)
	if err != nil {
		return err
	}

	ctx.(*gin.Context).SetCookie(RedisCookieKey, name, global.Config.WebAPP.Expires, "", global.Config.WebAPP.Domain, false, true)
	return nil
}

func GetKey(ctx context.Context, connName string, key string) (interface{}, error) {
	client, err := conn.GetConnect(connName)
	if err != nil {
		return nil, err
	}

	return client.Get(key).Result()
}

func SetKey(ctx context.Context, connName string, key string, value string) (string, error) {
	client, err := conn.GetConnect(connName)
	if err != nil {
		return "", err
	}
	return client.Set(key, value, -1).Result()
}

func Keys(ctx context.Context, connName string, pattern string) ([]string, error) {
	client, err := conn.GetConnect(connName)
	if err != nil {
		return nil, err
	}
	return client.Keys(pattern).Result()
}
