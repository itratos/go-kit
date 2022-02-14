package mongo

import (
	"context"
	"fmt"
	"log"

	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"gopkg.in/mgo.v2/bson"
)

var (
	NotFound = mongo.ErrNoDocuments.Error()
)

func (m *mgo) Insert(dbName, collectionName string, item interface{}) (string, error) {
	collection := m.client.Database(dbName).Collection(collectionName)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	res, err := collection.InsertOne(ctx, item)
	defer cancel()
	if err != nil {
		return "", err
	}
	return fmt.Sprint(res.InsertedID), nil
}

func (m *mgo) Get(dbName, collectionName string, id int) ([]byte, error) {
	collection := m.client.Database(dbName).Collection(collectionName)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	var result bson.M
	err := collection.FindOne(ctx, bson.M{"id": id}).Decode(&result)
	defer cancel()
	if err != nil {
		return nil, err
	}
	bsonBytes, erro := bson.Marshal(result)
	if erro != nil {
		return nil, erro
	}
	return bsonBytes, nil
}

func (m *mgo) GetOneByFiler(dbName, collectionName string, filter interface{}) ([]byte, error) {
	collection := m.client.Database(dbName).Collection(collectionName)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	var result bson.M
	err := collection.FindOne(ctx, filter).Decode(&result)
	defer cancel()
	if err != nil {
		return nil, err
	}
	bsonBytes, erro := bson.Marshal(result)
	if erro != nil {
		return nil, erro
	}
	return bsonBytes, nil
}

func (m *mgo) GetOneByFilterWd(dbname, collectionName string, filter interface{}) ([]byte, error) {
	collection := m.client.Database(dbname).Collection(collectionName)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	var result bson.M
	err := collection.FindOne(ctx, filter).Decode(&result)
	defer cancel()
	if err != nil {
		return nil, err
	}
	bsonBytes, erro := bson.Marshal(result)
	if erro != nil {
		return nil, erro
	}
	return bsonBytes, nil
}

func (m *mgo) GetByFilter(dbName, collectionName string, filter interface{}) ([][]byte, error) {
	log.Println(filter)
	collection := m.client.Database(dbName).Collection(collectionName)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	var results []bson.M
	cursor, err := collection.Find(ctx, filter)
	defer cancel()
	if err != nil {
		return nil, err
	}
	if err = cursor.All(ctx, &results); err != nil {
		panic(err)
	}
	var bsonBytesArray [][]byte
	for _, result := range results {
		bsonBytes, erro := bson.Marshal(result)
		if erro != nil {
			return nil, erro
		}
		bsonBytesArray = append(bsonBytesArray, bsonBytes)
	}

	return bsonBytesArray, nil
}

func (m *mgo) Update(dbName, collectionName string, item interface{}, id int64) (int, error) {

	collection := m.client.Database(dbName).Collection(collectionName)
	filter := bson.M{"id": id}
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	update := bson.M{
		"$set": item,
	}
	result, err := collection.UpdateOne(ctx, filter, update)
	defer cancel()
	if err != nil {
		return 0, err
	}
	return int(result.ModifiedCount), nil
}

func (m *mgo) UpdateWithFilter(dbName, collectionName string, item interface{}, filter interface{}) (int, error) {

	collection := m.client.Database(dbName).Collection(collectionName)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	update := bson.M{
		"$set": item,
	}
	result, err := collection.UpdateOne(ctx, filter, update)
	defer cancel()
	if err != nil {
		return 0, err
	}
	return int(result.ModifiedCount), nil
}

func (m *mgo) Delete(dbName, collectionName string, id int) (int, error) {

	collection := m.client.Database(dbName).Collection(collectionName)
	filter := bson.M{"id": id}
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	result, err := collection.DeleteOne(ctx, filter)
	defer cancel()
	if err != nil {
		return 0, err
	}
	return int(result.DeletedCount), nil
}

func (m *mgo) DeleteByFilter(dbName, collectionName string, filter interface{}) (int, error) {

	collection := m.client.Database(dbName).Collection(collectionName)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	result, err := collection.DeleteOne(ctx, filter)
	defer cancel()
	if err != nil {
		return 0, err
	}
	return int(result.DeletedCount), nil
}

func (m *mgo) GetId(dbName, collectionName string) int64 {
	collection := m.client.Database(dbName).Collection(collectionName)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	id, err := collection.CountDocuments(ctx, bson.M{})
	if err != nil {
		log.Println(err)
		return 0
	}
	return id + int64(1)
}
