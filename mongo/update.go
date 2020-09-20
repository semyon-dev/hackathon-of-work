package mongo

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"hackathon-work/model"
	"log"
	"strconv"
)

func UpdateStoreAnswers(stores model.CandidateStore, number, answer string) {
	filter := bson.M{"_id": stores.Id}
	var update bson.M
	switch number {
	case "q4":
		stores.Q4 = append(stores.Q4, answer)
		update = bson.M{"$set": bson.M{"q4": stores.Q4}}
	case "q6":
		stores.Q6 = append(stores.Q6, answer)
		update = bson.M{"$set": bson.M{"q6": stores.Q6}}
	case "q7":
		stores.Q7 = append(stores.Q7, answer)
		update = bson.M{"$set": bson.M{"q7": stores.Q7}}
	case "q8":
		stores.Q8 = append(stores.Q8, answer)
		update = bson.M{"$set": bson.M{"q8": stores.Q8}}
	case "q9":
		stores.Q9 = append(stores.Q9, answer)
		update = bson.M{"$set": bson.M{"q9": stores.Q9}}
	case "q11":
		r, err := strconv.Atoi(answer)
		if err != nil {
			//log.Println(err)
			return
		}
		stores.Q11 = r
		update = bson.M{"$set": bson.M{"q11": stores.Q11}}
	}
	_, err := candidatesStoreCollection.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		log.Fatal(err)
	}
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
