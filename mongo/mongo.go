package mongo

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var client *mongo.Client

func Connect() {

	var err error

	ctx := context.Background()

	client, err = mongo.Connect(ctx, options.Client().ApplyURI(
		"mongodb+srv://root:@url",
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
	//devDatabase := client.Database("dev")
	//fmt.Println("user collection name " + usersCollection.Name())
}
