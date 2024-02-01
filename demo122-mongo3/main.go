package main

import (
	"context"
	"fmt"
	"io/ioutil"
	"strings"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	// MongoDB 连接信息
	mongoConfig := map[string]interface{}{
		"user":                "root",
		"password":            "123456",
		"db_name":             "qskm",
		"addrs":               []string{"10.10.10.125:37017"},
		"replica_set":         "",
		"read_from_secondary": false,
	}

	// 创建 MongoDB 连接字符串
	connString := createMongoConnString(mongoConfig)

	// 连接 MongoDB
	client, err := mongo.NewClient(options.Client().ApplyURI(connString))
	if err != nil {
		fmt.Println("Error creating MongoDB client:", err)
		return
	}

	// 连接数据库
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	err = client.Connect(ctx)
	if err != nil {
		fmt.Println("Error connecting to MongoDB:", err)
		return
	}
	defer client.Disconnect(ctx)

	// 执行初始化语句
	initMongoTables(ctx, client, "/Users/me/Downloads/qskm_1.9.01.0.3.release_amd64/qskm/sql/init_mongo_table.sql")
}

func createMongoConnString(config map[string]interface{}) string {
	var connectionString string

	if user, ok := config["user"].(string); ok {
		connectionString += user + ":"
		if password, ok := config["password"].(string); ok {
			connectionString += password + "@"
		}
	}

	if addrs, ok := config["addrs"].([]string); ok {
		connectionString += strings.Join(addrs, ",")
	}

	if db, ok := config["db_name"].(string); ok {
		connectionString += "/" + db
	}

	return "mongodb://" + connectionString
}

func initMongoTables(ctx context.Context, client *mongo.Client, filePath string) {
	// 读取初始化语句文件内容
	content, err := ioutil.ReadFile(filePath)
	if err != nil {
		fmt.Println("Error reading initialization file:", err)
		return
	}

	// 将文件内容拆分为多个语句
	statements := strings.Split(string(content), ";")

	// 获取数据库
	dbName := client.Database("qskm")

	// 执行每个语句
	for _, statement := range statements {
		trimmedStatement := strings.TrimSpace(statement)
		if trimmedStatement != "" {
			err := dbName.RunCommand(ctx, map[string]interface{}{
				"eval": trimmedStatement,
			}).Err()
			if err != nil {
				fmt.Println("Error executing initialization statement:", err)
				return
			}
		}
	}

	fmt.Println("MongoDB initialization completed.")
}
