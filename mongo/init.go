package mongo

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Mgo interface {
	Insert(collectionName string, item interface{}) (string, error)
	Get(collectionName string, id int) ([]byte, error)
	GetOneByFiler(collectionName string, filter interface{}) ([]byte, error)
	GetByFilter(collectionName string, filter interface{}) ([][]byte, error)
	Update(collectionName string, item interface{}, id int64) (int, error)
	UpdateWithFilter(collectionName string, item interface{}, filter interface{}) (int, error)
	Delete(collectionName string, id int) (int, error)
	DeleteByFilter(collectionName string, filter interface{}) (int, error)
	GetId(collectionName string) int64
	GetOneByFilterWd(dbname, collectionName string, filter interface{}) ([]byte, error)
}

type mgo struct {
	client   *mongo.Client
	database string
}

func InitializeMongo(user, pass, database string) (Mgo, error) {
	clientOptions := options.Client().
		//ApplyURI("mongodb+srv://" + user + ":" + pass + "@dsms.tywcl.mongodb.net/dsms?retryWrites=true&w=majority")
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
