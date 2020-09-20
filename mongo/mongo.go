package mongo

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"hackathon-work/config"
	"hackathon-work/model"
	"log"
)

var client *mongo.Client
var candidatesStoreCollection *mongo.Collection
var candidatesRestaurantCollection *mongo.Collection
var candidatesDriversCollection *mongo.Collection

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
}

func GetAllRestaurants() (restaurants []model.CandidateRestaurant) {
	cursor, err := candidatesRestaurantCollection.Find(context.TODO(), bson.M{})
	if err != nil {
		log.Println(err)
	}
	if err = cursor.All(context.TODO(), &restaurants); err != nil {
		log.Println(err)
	}
	return restaurants
}

func GetAllStores() (cands []model.CandidateStore) {
	cursor, err := candidatesStoreCollection.Find(context.TODO(), bson.M{})
	if err != nil {
		log.Println(err)
	}
	if err = cursor.All(context.TODO(), &cands); err != nil {
		log.Println(err)
	}
	return cands
}
