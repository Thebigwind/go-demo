package main

import (
	"context"
	"fmt"
	"io/ioutil"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

func main() {
	// 连接到 MongoDB
	client, err := mongo.NewClient(options.Client().ApplyURI("your_mongo_uri"))
	if err != nil {
		log.Fatal(err)
	}
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}
	defer client.Disconnect(ctx)

	// Ping 数据库以确保连接成功
	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		log.Fatal(err)
	}

	// 读取并执行 init_mongo.js 脚本
	jsScript, err := ioutil.ReadFile("init_mongo.js")
	if err != nil {
		log.Fatal(err)
	}

	// 使用 MongoDB 的 eval 命令执行脚本
	result := client.Database("your_database").RunCommand(ctx, bson.D{{"eval", string(jsScript)}})
	if result.Err() != nil {
		// 在这里处理因为重复执行脚本导致的错误
		if result.Err().Error() == "your_specific_error_message" {
			// 忽略错误或做相应处理
			fmt.Println("Ignored error: ", result.Err())
		} else {
			log.Fatal(result.Err())
		}
	}

	// 其他操作...
}
