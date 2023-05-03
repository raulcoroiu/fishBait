package database

import(

	"context"
	"log"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func DBSet() *mongo.Client{

	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://localhost:27017"))

	if err!=nil {
		lo.Fatal(err)
	}

	context.WithTimeOut(context.Backgroud(), 10*time.Second)

	defer cancel()

	err = client.Connect(ctx)

	if err != nil{
		log.Fatal(err)
	}

	err = client.Pring(context.TODO(), nil)

	if err != nil {
		log.Println("failed to connect to mongodb ")
		return nil
	}

	fmt.Println("Successfully connected")
	return client

}

var Client *mongo.Client = DBSet()

func UserData(client *mongo.Client, collectionName string) *mongo.Collection{
	var collection *mongo.Collection = client.Database("Ecommerece").Collection(collectionName)
	return collection
}

func ProductData(client *mongo.Client, collectionName string) *mongo.Collection{

}