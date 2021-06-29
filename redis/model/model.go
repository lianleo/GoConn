package model

// Redis = redis.NewClient(&redis.Options{
// 	Addr:         config.Addr,
// 	ReadTimeout:  time.Duration(config.Timeout) * time.Second,
// 	WriteTimeout: time.Duration(config.Timeout) * time.Second,
// 	Password:     config.Password,
// 	DB:           config.DB,
// })
// _, err := Redis.Ping().Result()

type ConnConfig struct {
	Addr     string `json:"addr" bson:"addr"`
	Password string `json:"password" bson:"password"`
	DB       int    `json:"db" bson:"db"`
}

type RedisRequest struct {
	Command string `json:"commond"`
}
