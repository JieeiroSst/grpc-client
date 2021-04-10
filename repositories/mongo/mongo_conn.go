package mongo 

import (
	"context"
	"fmt"
	"sync"

	"github.com/JIeeiroSst/go-app/entities"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

var (
	mutex    sync.Mutex
	instance *Mongoconn
)

type Mongoconn struct {
	collection *mongo.Collection
}

type Config struct {
	DSN        string
	DB         string
	Collection string
}

func GetMongoConnInstance(cf Config) *Mongoconn {
	fmt.Println(cf)
	if instance == nil {
		mutex.Lock()
		defer mutex.Unlock()
		if instance == nil {
			dsn := cf.DSN
			clientOptions := options.Client().ApplyURI(dsn)
			client, err := mongo.Connect(context.TODO(), clientOptions)
			if err != nil {
				panic(err)
			}

			if err := client.Ping(context.TODO(), readpref.Primary()); err != nil {
				panic(err)
			}
			instance = &Mongoconn{
				collection: client.Database(cf.DB).Collection(cf.Collection),
			}
		}
	}
	return instance
}

func NewMongoSqlRepo(cf Config) *Mongoconn {
	return &Mongoconn{
		collection: GetMongoConnInstance(cf).collection,
	}
}

func (mongo *Mongoconn) CreateItem(item entities.Item) error {
	_,err:=mongo.collection.InsertOne(context.TODO(),item)
	return err
}

func (mongo *Mongoconn) FindAll() ([]entities.Item,error) {
	var items []entities.Item 
	cur, err := mongo.collection.Find(context.TODO(), bson.D{{}})
	if err != nil {
		return nil, err
	}
	for cur.Next(context.TODO()) {
		var item entities.Item
		_ = cur.Decode(&item)
	}
	return items, nil
}

func (mongo *Mongoconn) FindByID(id string) (entities.Item,error) {
	filter := bson.D{{"_id", id}}
	result := entities.Item{}
	err := mongo.collection.FindOne(context.TODO(), filter).Decode(&result)
	if err != nil {
		return entities.Item{}, err
	}
	return result, nil
}

func (mongo *Mongoconn) DeleteItem(id string) error{
	filter := bson.D{{"_id", id}}
	_, err := mongo.collection.DeleteOne(context.TODO(), filter)
	if err != nil {
		return err
	}
	return nil
}

func (mongo *Mongoconn) UpdateItem(item entities.Item,id string) error {
	filter := bson.M{"_id": bson.M{"$eq": id}}
	update := bson.D{
		{"$set", bson.D{{"quantity", item.Quantity}}},
	}
	_, err := mongo.collection.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		return err
	}
	return nil
}