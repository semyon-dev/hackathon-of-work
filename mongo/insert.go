package mongo

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"hackathon-work/model"
	"log"
)

func InsertRestaurantAnswers(answers model.CandidateRestaurant) (err error) {
	_, err = candidatesRestaurantCollection.InsertOne(context.Background(), answers)
	return err
}

func InsertStoreAnswers(answers model.CandidateStore) (err error) {
	_, err = candidatesStoreCollection.InsertOne(context.Background(), answers)
	return err
}

func UpdateStoreAnswers(answers model.CandidateStore) (err error) {
	filter := bson.M{"_id": answers.Id}
	update := bson.M{"$set": bson.M{"desciption": answers.Description}}
	_, err = candidatesStoreCollection.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		log.Println(err)
	}
	return err
}

func UpdateRestaurantsAnswers(answers model.CandidateRestaurant, number, answer string) {
	filter := bson.M{"_id": answers.Id}
	var update bson.M
	switch number {
	case "q3":
		answers.Q3 = answer
		update = bson.M{"$set": bson.M{"q3": answers.Q3}}
	case "q6":
		answers.Q6 = append(answers.Q6, answer)
		update = bson.M{"$set": bson.M{"q6": answers.Q6}}
	case "q7":
		answers.Q7 = append(answers.Q7, answer)
		update = bson.M{"$set": bson.M{"q7": answers.Q7}}
	}
	_, err := candidatesRestaurantCollection.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		log.Fatal(err)
	}
}

func UpdatesDriversAnswers(answers model.CandidateDriver) (err error) {
	filter := bson.M{"_id": answers.Id}
	update := bson.M{"$set": bson.M{"desciption": answers.Description}}
	_, err = candidatesDriversCollection.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		log.Println(err)
	}
	return err
}

func InsertDriversAnswers(answers model.CandidateDriver) (err error) {
	_, err = candidatesDriversCollection.InsertOne(context.Background(), answers)
	return err
}
