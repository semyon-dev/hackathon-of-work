package mongo

import (
	"context"
	"hackathon-work/model"
)

func InsertRestaurantAnswers(answers model.CandidateRestaurant) (err error) {
	_, err = candidatesRestaurantCollection.InsertOne(context.Background(), answers)
	return err
}

func InsertStoreAnswers(answers model.CandidateStore) (err error) {
	_, err = candidatesStoreCollection.InsertOne(context.Background(), answers)
	return err
}

func InsertDriversAnswers(answers model.CandidateDriver) (err error) {
	_, err = candidatesDriversCollection.InsertOne(context.Background(), answers)
	return err
}


