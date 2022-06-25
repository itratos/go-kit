package mongo

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Mgo interface {
	Insert(dbName, collectionName string, item interface{}) (string, error)
	Get(dbName, collectionName string, id int) ([]byte, error)
	GetByUuid(dbName, collectionName string, id string) ([]byte, error)
	GetOneByFiler(dbName, collectionName string, filter interface{}) ([]byte, error)
	GetByFilter(dbName, collectionName string, filter interface{}) ([][]byte, error)
	Update(dbName, collectionName string, item interface{}, id int64) (int, error)
	UpdateByUuid(dbName, collectionName string, item interface{}, id string) (int, error)
	UpdateWithFilter(dbName, collectionName string, item interface{}, filter interface{}) (int, error)
	Delete(dbName, collectionName string, id int) (int, error)
	DeleteByUuid(dbName, collectionName string, id string) (int, error)
	DeleteByFilter(dbName, collectionName string, filter interface{}) (int, error)
	GetId(dbName, collectionName string) int64
	GetClient() *mongo.Client
}

type mgo struct {
	client   *mongo.Client
	database string
}

func InitializeMongo(user, pass, database string) (Mgo, error) {
	clientOptions := options.Client().
		ApplyURI("mongodb+srv://" + user + ":" + pass + "@" + database + ".tywcl.mongodb.net/" + database + "?retryWrites=true&w=majority")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	mongo, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	log.Println("mongoDb connected successfully")
	cliente := &mgo{
		client:   mongo,
		database: database,
	}
	return cliente, nil
}
