package main

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"strings"
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
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	err = client.Connect(ctx)
	if err != nil {
		fmt.Println("Error connecting to MongoDB:", err)
		return
	}
	defer client.Disconnect(ctx)

	// 执行初始化脚本
	initMongoTables(ctx, client, `
		db.createCollection("test_lxf");
		db.test_lxf.createIndex({"metadata.id": 1});
		db.createCollection("test_luxuefeng");
		db.test_luxuefeng.createIndex({"metadata.id": 1});
	`)
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

func initMongoTables(ctx context.Context, client *mongo.Client, script string) {
	statements := strings.Split(script, ";")

	for _, statement := range statements {
		trimmedStatement := strings.TrimSpace(statement)
		if trimmedStatement != "" {
			err := client.Database("qskm").RunCommand(ctx, map[string]interface{}{
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
