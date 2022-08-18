package main

import (
	"context"
	"fmt"
	"time"

	"github.com/labstack/gommon/log"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// 消息盒子
type NoticeTest struct {
	Type       int    `bson:"type" json:"type"`
	UserId     int    `bson:"user_id" json:"user_id"`
	Title      string `bson:"title" json:"title"`
	Content    string `bson:"content" json:"content"`
	ReadStatus int    `bson:"read_status" json:"read_status"`
	CreateTime int64  `bson:"create_time" json:"create_time"`
	UpdateTime int64  `bson:"update_time" json:"update_time"`
}

func GetMgo() *mongo.Client {
	// 超时设置
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// 连接MongoDB
	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://admin:admin_1xm@192.168.1.16:27017"))
	if err != nil {
		log.Fatal("Connect-Err:", err)
	}

	// 检查连接
	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatal("Ping-Err:", err)
	}

	return client
}

func OptCollection(db, name string) *mongo.Collection {
	return GetMgo().Database(db).Collection(name)
}

// 新增文档
func Insert(n NoticeTest) {
	collection := OptCollection("coinsky_app", "cs_notice_test")
	result, err := collection.InsertOne(context.TODO(), n)
	if err != nil {
		log.Fatal("Insert-Err:", err)
	}

	fmt.Println("新增单个文档：", result.InsertedID)
}

// 新增多文档
func InsertAll(n []interface{}) {
	collection := OptCollection("coinsky_app", "cs_notice_test")
	result, err := collection.InsertMany(context.TODO(), n)
	if err != nil {
		log.Fatal("InsertMany-Err:", err)
	}

	fmt.Println("新增多个文档：", result.InsertedIDs)
}

// 删除文档
func Delete(filter interface{}) {
	collection := OptCollection("coinsky_app", "cs_notice_test")
	result, err := collection.DeleteOne(context.TODO(), filter)
	if err != nil {
		log.Fatal("Delete-Err:", err)
	}

	fmt.Println("删除单个文档：", result.DeletedCount)
}

// 删除多文档
func DeleteAll(filter interface{}) {
	collection := OptCollection("coinsky_app", "cs_notice_test")
	result, err := collection.DeleteMany(context.TODO(), filter)
	if err != nil {
		log.Fatal("DeleteMany-Err:", err)
	}

	fmt.Println("删除多个文档：", result.DeletedCount)
}

// 更新文档
func Update(filter, update interface{}) {
	collection := OptCollection("coinsky_app", "cs_notice_test")
	result, err := collection.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		log.Fatal("Update-Err:", err)
	}

	fmt.Println("更新单个文档：", result.MatchedCount, result.ModifiedCount)
}

// 更新多个文档
func UpdateAll(filter, update interface{}) {
	collection := OptCollection("coinsky_app", "cs_notice_test")
	result, err := collection.UpdateMany(context.TODO(), filter, update)
	if err != nil {
		log.Fatal("UpdateMany-Err:", err)
	}

	fmt.Println("更新多文档：", result.MatchedCount, result.ModifiedCount)
}

// 查询
func Find(filter interface{}) {
	var notice NoticeTest

	collection := OptCollection("coinsky_app", "cs_notice_test")
	err := collection.FindOne(context.TODO(), filter).Decode(&notice)
	if err != nil {
		log.Fatal("Find-Err:", err)
	}

	fmt.Println("查询单文档：", notice)
}

// 查询多个文档
func FindAll(filter interface{}) {
	// 定义切片存储查询结果
	var noticeList []*NoticeTest

	findOptions := options.Find()
	findOptions.SetLimit(10)

	collection := OptCollection("coinsky_app", "cs_notice_test")
	cur, err := collection.Find(context.TODO(), filter, findOptions)
	if err != nil {
		log.Fatal("FindAll-Err:", err)
	}

	// 查找多个文档返回一个光标
	// 遍历游标允许解码文档
	for cur.Next(context.TODO()) {
		// 创建一个值，将单个文档解码
		var elem NoticeTest
		err := cur.Decode(&elem)
		if err != nil {
			log.Fatal("Find-Decode:", err)
		}
		noticeList = append(noticeList, &elem)
	}

	if err := cur.Err(); err != nil {
		log.Fatal("Find-Cursor-Err:", err)
	}

	cur.Close(context.TODO())
	fmt.Printf("查询多个文档：%#v\n", noticeList)
}

func FindAll2(filter interface{}) {
	findOptions := options.Find()
	findOptions.SetLimit(10)

	collection := OptCollection("coinsky_app", "cs_notice_test")
	cur, err := collection.Find(context.TODO(), filter, findOptions)
	if err != nil {
		log.Fatal("FindAll-Err:", err)
	}

	var resultsM []bson.M
	if err := cur.All(context.TODO(), &resultsM); err != nil {
		log.Fatal("Cursor-All-Err:", err)
	}

	// 遍历的顺序是随机的
	for _, result := range resultsM {
		for key, value := range result {
			fmt.Println("字段=", key, "  值=", value)
		}
	}

	fmt.Println("查询多个文档：", resultsM)
}

func FindAll3(filter interface{}) {
	findOptions := options.Find()

	results := make([]NoticeTest, 0)
	collection := OptCollection("coinsky_app", "cs_notice_test")
	cur, err := collection.Find(context.TODO(), filter, findOptions)
	if err = cur.Err(); err != nil {
		log.Fatal(err)
	}

	for cur.Next(context.TODO()) {
		var elem NoticeTest
		err := cur.Decode(&elem)
		if err != nil {
			log.Fatal(err)
		}

		results = append(results, elem)
	}

	fmt.Println("查询多个文档：", results)
}

func main() {
	//notice := NoticeTest{
	//	Type:       1,
	//	UserId:     10023,
	//	Title:      "Nft Tracker: Moonfrogs Official最近1小时发生698次Mints，趋势十分火爆。",
	//	Content:    "Nft Tracker: Moonfrogs Official最近1小时发生698次Mints，趋势十分火爆。",
	//	ReadStatus: 0,
	//	CreateTime: time.Now().Unix(),
	//	UpdateTime: time.Now().Unix(),
	//}
	//Insert(notice)

	//notice1 := NoticeTest{
	//	Type:       1,
	//	UserId:     10021,
	//	Title:      "Nft Tracker: Random AI Official最近1小时发生700次Mints，趋势十分火爆。",
	//	Content:    "Nft Tracker: Random AI Official最近1小时发生700次Mints，趋势十分火爆。",
	//	ReadStatus: 0,
	//	CreateTime: time.Now().Unix(),
	//	UpdateTime: time.Now().Unix(),
	//}
	//notice2 := NoticeTest{
	//	Type:       2,
	//	UserId:     10022,
	//	Title:      "Nft Tracker: Random AI Official最近1小时发生700次Mints，趋势十分火爆。",
	//	Content:    "Nft Tracker: Random AI Official最近1小时发生700次Mints，趋势十分火爆。",
	//	ReadStatus: 0,
	//	CreateTime: time.Now().Unix(),
	//	UpdateTime: time.Now().Unix(),
	//}
	//notice3 := NoticeTest{
	//	Type:       3,
	//	UserId:     10023,
	//	Title:      "Nft Tracker: Random AI Official最近1小时发生700次Mints，趋势十分火爆。",
	//	Content:    "Nft Tracker: Random AI Official最近1小时发生700次Mints，趋势十分火爆。",
	//	ReadStatus: 0,
	//	CreateTime: time.Now().Unix(),
	//	UpdateTime: time.Now().Unix(),
	//}
	//noticeList := []interface{}{notice1, notice2, notice3}
	//InsertAll(noticeList)

	//filter := bson.D{{"type", 1}}
	//Delete(filter)
	//DeleteAll(filter)

	//filter := bson.D{{"type", 3}}
	//update := bson.D{
	//	{"$set", bson.D{
	//		{"title", "更新测试..."},
	//		{"read_status", 1},
	//	}},
	//}
	//Update(filter, update)
	//UpdateAll(filter, update)

	//filter := bson.D{{"type", 2}}
	//Find(filter)
	//FindAll(filter)
	//FindAll2(filter)
	//FindAll3(filter)
}
