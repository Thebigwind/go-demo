package main

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"log"
	"time"
)

type MongoDB struct {
	Host        string `yaml:"host"`
	Port        string `yaml:"port"`
	User        string `yaml:"user"`
	Password    string `yaml:"password"`
	ConnTimeOut int    `yaml:"conn_time_out"`
	// MaxOpenConns int    `yaml:"max_open_conns"`
	DbName string `yaml:"db_name"`
}

var mongoDB *mongo.Database

func InitMongoDB(conf MongoDB) {
	url := ""
	if conf.Password == "" {
		url = fmt.Sprintf("mongodb://%s:%s/?retryWrites=false", conf.Host, conf.Port)
	} else {
		//mongodb://[username:password@]host1[:port1][,host2[:port2],...[,hostN[:portN]]][/[database][?options]]
		url = fmt.Sprintf("mongodb://%s:%s@%s:%s/?retryWrites=false", conf.User, conf.Password, conf.Host, conf.Port)
	}
	client, err := mongo.NewClient(options.Client().ApplyURI(url))
	if err != nil {
		fmt.Printf("err:%v", err)
	}
	//ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	ctx := context.Background()
	errC := client.Connect(ctx)
	if errC != nil {
		fmt.Printf("err:%v", err)
	}
	errP := client.Ping(ctx, readpref.Primary())
	if errP != nil {
		fmt.Printf("err:%v", err)
	}

	session, _ := client.StartSession()
	err = session.StartTransaction()
	if err != nil {
		fmt.Printf("err:%v", err)
	}
	mongoDB = client.Database(conf.DbName)
}

func GetMongoDB() *mongo.Database {
	return mongoDB
}

func main() {
	conf := MongoDB{
		Host:        "10.10.10.162",
		Port:        "37027",
		User:        "root",
		Password:    "123456",
		ConnTimeOut: 30,
		DbName:      "audit",
	}
	InitMongoDB(conf)
}

func test() {
	// 设置MongoDB连接选项
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")

	// 创建MongoDB客户端
	client, err := mongo.NewClient(clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	// 连接MongoDB
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}
	defer client.Disconnect(ctx)

	// 尝试访问一个MongoDB集合来检查连接
	collection := client.Database("your_database_name").Collection("your_collection_name")

	// 通过执行FindOne操作来检查连接
	var result bson.M
	err = collection.FindOne(ctx, bson.M{}).Decode(&result)
	if err != nil {
		log.Fatalf("无法访问MongoDB集合：%v", err)
	}

	fmt.Println("MongoDB连接正常，可以正常访问MongoDB集合。")
}
