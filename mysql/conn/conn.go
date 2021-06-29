package conn

import (
	"database/sql"
	"fmt"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/lianleo/GoConn/mysql/model"
)

var dbMap map[string]*sql.DB

//init 初始化连接Map
func init() {
	dbMap = make(map[string]*sql.DB)
}

//Connect 连接Redis，创建Redis客户端
func Connect(name string, cfg model.ConnConfig) error {
	url := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", cfg.UserName, cfg.Password, cfg.IP, cfg.Port, cfg.Database)
	DB, err := sql.Open("mysql", url)
	if err != nil {
		fmt.Printf("Open mysql failed,err:%v\n", err)
		return err
	}
	DB.SetConnMaxLifetime(100 * time.Second) //最大连接周期，超过时间的连接就close
	DB.SetMaxOpenConns(10)                   //设置最大连接数
	DB.SetMaxIdleConns(3)                    //设置闲置连接数
	dbMap[name] = DB
	return nil
}

//GetConnect 获取数据库连接客户端
func GetConnect(name string) (*sql.DB, error) {
	db, ok := dbMap[name]
	if ok {
		return db, nil
	}
	return nil, fmt.Errorf("%s连接不存在", name)
}
