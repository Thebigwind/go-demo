package main

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
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
