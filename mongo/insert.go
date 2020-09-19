package mongo

import (
	"context"
	"hackathon-work/model"
)

func InsertRestaurantAnswers(answers model.CandidateStore) (err error) {
	_, err = candidatesRestaurantCollection.InsertOne(context.Background(), answers)
	return err
}


