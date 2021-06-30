package model

type ConnConfig struct {
	UserName string `json:"username" bson:"username" toml:"username"`
	Password string `json:"password" bson:"password" toml:"password"`
	IP       string `json:"ip" bson:"ip" toml:"ip"`
	Port     int    `json:"port" bson:"port" toml:"port"`
	Database string `json:"database" bson:"database" toml:"database"`
}
