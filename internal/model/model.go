package model

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"time"

	"spreadsheet-server/configs"
)

// MongoDBClient MongoDB 连接池
var MongoDBClient *mongo.Database

// ConnectDB 初始化数据库链接
func ConnectDB() {

	fmt.Printf("appconfig %+v", configs.AppConfig)
	mongodbConf := configs.AppConfig.Db["mongodb"]

	user := mongodbConf.User
	password := mongodbConf.Password
	host := mongodbConf.Host
	port := mongodbConf.Port
	dbName := mongodbConf.Database
	timeOut := 2
	maxNum := 50

	uri := fmt.Sprintf("mongodb://%s:%s@%s:%s/%s?w=majority", user, password, host, port, dbName)
	// 设置连接超时时间
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(timeOut))
	defer cancel()
	// 通过传进来的uri连接相关的配置
	o := options.Client().ApplyURI(uri)
	// 设置最大连接数 - 默认是100 ，不设置就是最大 max 64
	o.SetMaxPoolSize(uint64(maxNum))
	// 发起链接
	client, err := mongo.Connect(ctx, o)
	if err != nil {
		fmt.Println("ConnectToDB", err)
		return
	}
	// 判断服务是不是可用
	if err = client.Ping(context.Background(), readpref.Primary()); err != nil {
		fmt.Println("ConnectToDB", err)
		return
	}
	// 返回 client
	MongoDBClient = client.Database(dbName)
}
