package mongo

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"hackathon-work/model"
	"log"
)

// получение всех кандидатов на официантов
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
