package service

import (
	"context"
	"database/sql"
	"fmt"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/lianleo/GoCommon/log"
	"github.com/lianleo/GoConn/global"
	"github.com/lianleo/GoConn/mysql/conn"
	"github.com/lianleo/GoConn/mysql/model"
)

const (
	MysqlCookieKey = "CurrentMysqlConnection"
)

func Connect(ctx context.Context, name string, config model.ConnConfig) error {
	err := conn.Connect(name, config)
	if err != nil {
		return err
	}

	ctx.(*gin.Context).SetCookie(MysqlCookieKey, name, global.Config.WebAPP.Expires, "", global.Config.WebAPP.Domain, false, true)
	return nil
}

func Insert(ctx context.Context, name string, data interface{}) error {

	f := func(idx string) {
		begin := time.Now()

		db, err := conn.GetConnect(name)
		if err != nil {
			log.Error(err)
		}

		// tx, err := db.Begin()
		tx, err := db.BeginTx(ctx, nil)
		if err != nil {
			log.Error(err)
		}
		//准备sql语句
		stmt, err := tx.Prepare("INSERT INTO sys_user (`id`, `name`, `phone`) VALUES (?, ?, ?)")
		if err != nil {
			fmt.Println("Prepare fail")
			log.Error(err)
		}
		//将参数传递到sql语句中并且执行
		for i := 0; i < 10*1; i++ {
			id := uuid.NewString()
			_, err := stmt.Exec(id, "leo", fmt.Sprintf("186%0*d", 8, i))
			if err != nil {
				fmt.Println("Exec fail")
				log.Error(err)
			}
			if i%100 == 0 {
				fmt.Printf("%s %d\n", idx, i)
			}
		}
		//将事务提交

		err = tx.Commit()
		if err != nil {
			fmt.Println("Commit fail")
			log.Error(err)
		}

		end := time.Now()

		fmt.Printf("开始 %s::%d \n", idx, begin.UnixNano())
		fmt.Printf("结束 %s::%d \n", idx, end.UnixNano())
		fmt.Printf("用时 %s::%d \n", idx, (end.UnixNano()-begin.UnixNano())/int64(time.Millisecond))
		fmt.Printf("用时 %s::%d \n", idx, (end.Unix() - begin.Unix()))
	}

	for i := 0; i < 1; i++ {
		go f(fmt.Sprintf("a%d ", i))
	}

	return nil
}

//RunSQL
func RunSQL(ctx context.Context, name, sqlang string) (rs interface{}, err error) {
	db, err := conn.GetConnect(name)
	if err != nil {
		log.Error(err)
	}
	if strings.Index(sqlang, "insert") == 0 || strings.Index(sqlang, "INSERT") == 0 || strings.Index(sqlang, "update") == 0 || strings.Index(sqlang, "UPDATE") == 0 {
		result, err := db.Exec(sqlang)
		if err != nil {
			log.Error(err)
		}
		log.Info(result)
		return &result, err
	}
	if strings.Index(sqlang, "select") == 0 {
		rows, err := db.Query(sqlang)

		defer func() {
			if rows != nil {
				rows.Close()
			}
		}()
		if err != nil {
			fmt.Printf("Query failed,err:%v", err)
			return nil, err
		}
		// type User struct {
		// 	ID    string
		// 	Name  string
		// 	Phone string
		// }

		cols, _ := rows.Columns()
		log.Info(cols)
		cts, _ := rows.ColumnTypes()
		for _, ct := range cts {
			log.Infof("%v", ct)
			l, _ := ct.Length()
			p, s, _ := ct.DecimalSize()
			log.Infof("%s %v length:%v precision:%v scale:%v", ct.Name(), ct.DatabaseTypeName(), l, p, s)
		}
		log.Info(cts)

		list := [](map[string]interface{}){}

		for rows.Next() {
			data := make(map[string]interface{}, len(cols))
			dl := []interface{}{}
			for _, ct := range cts {
				if ct.DatabaseTypeName() == "VARCHAR" {
					var v *string
					data[ct.Name()] = &v
					dl = append(dl, &v)
				} else if ct.DatabaseTypeName() == "INT" {
					var v *int64
					data[ct.Name()] = &v
					dl = append(dl, &v)
				} else {
					var v sql.NullString
					data[ct.Name()] = &v.String
					dl = append(dl, &v)
				}
			}

			err = rows.Scan(dl...)
			if err != nil {
				fmt.Printf("Scan failed,err:%v", err)
				return nil, err
			}
			list = append(list, data)
		}
		return list, nil
	}

	return nil, err
}

func OrmTest(ctx context.Context, name string, query string) (interface{}, error) {

	return nil, nil
}
