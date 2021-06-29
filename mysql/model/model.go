package model

type ConnConfig struct {
	UserName string `json:"username" bson:"username"`
	Password string `json:"password" bson:"password"`
	IP       string `json:"ip" bson:"ip"`
	Port     int    `json:"port" bson:"port"`
	Database string `json:"database" bson:"database"`
}
