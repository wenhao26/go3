package main

import (
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/bson"

	"coinsky_go_project/common/mongodb"
)

type CsNotice struct {
	UserId  int    `bson:"user_id"`
	Type    int    `bson:"type"`
	Title   string `bson:"title"`
	Content string `bson:"content"`
}

func main() {
	user := "admin"
	password := "admin_1xm"
	host := "192.168.1.16"
	port := "27017"
	dbName := "coinsky_app"
	timeout := 2
	poolSize := 50
	uri := fmt.Sprintf("mongodb://%s:%s@%s:%s", user, password, host, port)

	db, err := mongodb.ConnectDB(uri, dbName, time.Duration(timeout), uint64(poolSize))
	if err != nil {
		fmt.Println("初始化连接失败:", err)
		return
	}

	var notice CsNotice
	collection := db.SetCollection("cs_notice")
	_ = collection.FindOne(context.TODO(), bson.M{}).Decode(&notice)
	fmt.Println(notice)
}
