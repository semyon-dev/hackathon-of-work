package db

import (
	"fmt"
	"hackathon-work/model"
	"hackathon-work/mongo"
	"strconv"
)

func GetStoreAnswers() (err error) {

	candidates := GetCandidates(10, 0)

	for _, c := range candidates {

		var answers model.CandidateStore

		if c.Type == "store" {
			answers.Q1 = c.Position
			answers.Q2 = true

			temp, _ := strconv.Atoi(c.Experience)
			answers.Q3 = int64(temp)
			answers.Q4 = c.Industry
			answers.Q5 = c.Companies
			answers.Q6 = c.Systems
			answers.Q7 = c.ExperiencePrograms
			answers.Q8 = c.ExperienceTools
			answers.Q9 = c.TypesWork
			//answers.Q10 = c.FindWorkTime
			//answers.Q11 = c.Salary

			err := mongo.InsertStoreAnswers(answers)
			if err != nil {
				fmt.Println(err)
			}
		}
	}

	return err
}
