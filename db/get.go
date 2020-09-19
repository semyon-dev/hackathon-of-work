package db

import (
	"hackathon-work/model"
	"strconv"
)

func InsertStoreAnswers() (err error) {

	candidates := GetCandidates(10,0)

	for _, c := range candidates {

		var answers model.CandidateStore

		if c.Type == "store" {
			answers.Q1 = c.Position
			answers.Q2 = true

			temp, _ := strconv.Atoi(c.Experience)
			answers.Q3 = int64(temp)
			answers.Q4 = c.Industry
			//answers.Q5 = c.Companies
			answers.Q6 = c.Systems
			answers.Q7 = c.ExperiencePrograms
			//answers.Q8 = c.ExperienceTools
			//answers.Q9 = c.WorkTypes
			//answers.Q10 = c.FindWorkTime
			//answers.Q11 = c.Salary
		}
	}

	return err
}
