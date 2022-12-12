package main

import (
	"context"
	"log"
)

// 查询单个文档
func (m *mgo) FindOne(key string, value interface{}) *mongo.SingleResult {
	client := DB.Mongo
	collection, e := client.Database(m.database).Collection(m.collection).Clone()
	if e != nil {
		log.Fatal(e)
	}
	filter := bson.D{{key, value}}
	singleResult := collection.FindOne(context.TODO(), filter)
	return singleResult
}

//查询多个文档
func (m *mgo) FindMany(filter interface{}) (*mongo.Cursor, error) {
	client := DB.Mongo
	collection, e := client.Database(m.database).Collection(m.collection).Clone()
	if e != nil {
		log.Fatal(e)
	}
	return collection.Find(context.TODO(), filter)
}

//多条件查询
func (m *mgo) FindManyByFilters(filter interface{}) (*mongo.Cursor, error) {
	client := DB.Mongo
	collection, e := client.Database(m.database).Collection(m.collection).Clone()
	if e != nil {
		log.Fatal(e)
	}
	return collection.Find(context.TODO(), bson.M{"$and": filter})
}

//查询集合里有多少数据
func (m *mgo) CollectionCount() (string, int64) {
	client := DB.Mongo
	collection := client.Database(m.database).Collection(m.collection)
	name := collection.Name()
	size, _ := collection.EstimatedDocumentCount(context.TODO())
	return name, size
}

//按选项查询集合
// Skip 跳过
// Limit 读取数量
// sort 1 ，-1 . 1 为升序 ， -1 为降序
func (m *mgo) CollectionDocuments(Skip, Limit int64, sort int, key string, value interface{}) *mongo.Cursor {
	client := DB.Mongo
	collection := client.Database(m.database).Collection(m.collection)
	SORT := bson.D{{"_id", sort}}
	filter := bson.D{{key, value}}
	findOptions := options.Find().SetSort(SORT).SetLimit(Limit).SetSkip(Skip)
	temp, _ := collection.Find(context.Background(), filter, findOptions)
	return temp
}
