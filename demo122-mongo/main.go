package main

import (
	"context"
	"fmt"
	"io/ioutil"
	"log"
	"strings"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// ExecuteMongoCommands 用于执行 MongoDB 相关的命令
func ExecuteMongoCommands(client *mongo.Client, commands []string) error {
	for _, command := range commands {
		command = strings.TrimSpace(command)

		if len(command) > 0 {
			err := client.Database("qskm").RunCommand(context.TODO(), map[string]interface{}{
				"eval": command,
			}).Err()
			if err != nil {
				return err
			}
		}
	}

	return nil
}

func main() {
	// 连接到 MongoDB
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI("mongodb://10.10.10.125:37017"))
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		if err = client.Disconnect(context.TODO()); err != nil {
			log.Fatal(err)
		}
	}()

	// 读取 MongoDB 相关的命令文件
	filePath := "/Users/me/Downloads/qskm_1.9.01.0.3.release_amd64/qskm/sql/init_mongo_table.sql"
	content, err := ioutil.ReadFile(filePath)
	if err != nil {
		log.Fatal(err)
	}

	// 以分号作为分隔符拆分 MongoDB 命令
	commands := strings.Split(string(content), ";")

	// 执行 MongoDB 相关的命令
	err = ExecuteMongoCommands(client, commands)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("MongoDB commands executed successfully")
}
