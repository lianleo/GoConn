package service

import (
	"context"

	"github.com/globalsign/mgo/bson"
	"github.com/lianleo/GoConn/mongo/conn"
)

func AddCollection(ctx context.Context, database string, collectionName string) error {
	db, err := conn.GetConnect(database)
	if err != nil {
		return err
	}
	r := struct {
		ID    bson.ObjectId `bson:"_id"`
		Value string        `bson:"value"`
	}{
		bson.NewObjectId(),
		"hello world",
	}
	return db.C(collectionName).Insert(r)
}

func Insert(ctx context.Context, database string, coll string, data bson.M) error {
	db, err := conn.GetConnect(database)
	if err != nil {
		return err
	}
	return db.C(coll).Insert(data)
}

func Query(ctx context.Context, database string, coll string, query bson.M) (data []interface{}, err error) {
	db, err := conn.GetConnect(database)
	if err != nil {
		return nil, err
	}
	err = db.C(coll).Find(query).All(&data)
	return
}
