package mongo

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"hackathon-work/config"
)

var client *mongo.Client
var candidatesStoreCollection *mongo.Collection
var candidatesRestaurantCollection *mongo.Collection
var candidatesDriversCollection *mongo.Collection
var vacationsCollection *mongo.Collection

func Connect() {

	var err error

	ctx := context.Background()

	client, err = mongo.Connect(ctx, options.Client().ApplyURI(
		"mongodb+srv://root:"+config.MongoDBPassword+"@cluster0.qjp6o.gcp.mongodb.net/main?retryWrites=true&w=majority",
	))
	if err != nil {
		fmt.Println(err)
	}

	// Create connect
	err = client.Connect(context.TODO())
	if err != nil {
		fmt.Println("client MongoDB:", err)
	} else {
		fmt.Println("✔ Подключение client MongoDB успешно")
	}
	mainDB := client.Database("main")
	candidatesStoreCollection = mainDB.Collection("candidates_store")
	candidatesRestaurantCollection = mainDB.Collection("candidates_restaurant")
	candidatesDriversCollection = mainDB.Collection("candidates_drivers")
	vacationsCollection = mainDB.Collection("vacations")
}
