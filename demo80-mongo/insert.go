package main

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
)

//插入单个文档
func (m *mgo) InsertOne(value interface{}) *mongo.InsertOneResult {
	client := DB.Mongo
	collection := client.Database(m.database).Collection(m.collection)
	insertResult, err := collection.InsertOne(context.TODO(), value)
	if err != nil {
		log.Fatal(err)
	}
	return insertResult
}

//插入多个文档
func (m *mgo) InsertMany(values []interface{}) *mongo.InsertManyResult {
	client := DB.Mongo
	collection := client.Database(m.database).Collection(m.collection)
	result, err := collection.InsertMany(context.TODO(), values)
	if err != nil {
		log.Fatal(err)
	}
	return result
}
