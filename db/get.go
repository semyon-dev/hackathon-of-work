package db

import (
	"hackathon-work/model"
	"hackathon-work/mongo"
	"log"
	"strconv"
	"strings"
)

func GetStoreAnswers() {

	candidates := GetCandidates(3, 0)

	for _, c := range candidates {

		var answers model.CandidateStore

		if c.Type == "store" {
			answers.Q1 = c.Position
			answers.Q1 = c.Position
			answers.Q2 = true

			temp, err := strconv.Atoi(c.Experience)
			if err != nil {
				log.Println(err)
			}
			answers.Q3 = int64(temp)
			answers.Id = c.Id
			answers.CandidateID = c.CandidateID

			c.Industry = strings.Trim(c.Industry, "{}")
			c.Companies = strings.Trim(c.Companies, "{}")
			c.Systems = strings.Trim(c.Systems, "{}")
			c.ExperiencePrograms = strings.Trim(c.Systems, "{}")
			c.ExperienceTools = strings.Trim(c.ExperienceTools, "{}")
			c.TypesWork = strings.Trim(c.TypesWork, "{}")

			c.Industry = strings.Trim(c.Industry, "\"")
			c.Companies = strings.Trim(c.Companies, "\"")
			c.Systems = strings.Trim(c.Systems, "\"")
			c.ExperiencePrograms = strings.Trim(c.ExperiencePrograms, "\"")
			c.ExperienceTools = strings.Trim(c.ExperienceTools, "\"")
			c.TypesWork = strings.Trim(c.TypesWork, "\"")

			if c.TypesWork == "" && c.ExperienceTools == "" && c.Systems == "" && c.ExperiencePrograms == "" && c.Companies == "" && c.Industry == "" {
				continue
			}

			answers.Q4 = strings.Split(c.Industry, ",")
			answers.Q5 = strings.Split(c.Companies, ",")
			answers.Q6 = strings.Split(c.Systems, ",")
			answers.Q7 = strings.Split(c.ExperiencePrograms, ",")
			answers.Q8 = strings.Split(c.ExperienceTools, ",")
			answers.Q9 = strings.Split(c.TypesWork, ",")

			//answers.Q10 = c.FindWorkTime
			//answers.Q11 = c.Salary

			err = mongo.InsertStoreAnswers(answers)
			if err != nil {
				log.Println(err)
			}
		}
	}
}

// for updates
const limit = 10000
const offset = 90000

//func UpdateStoreAnswers(done chan struct{}) {
//	candidates := GetCandidates(limit, offset)
//	for _, c := range candidates {
//		var answers model.CandidateStore
//		if c.Type == "store" {
//			answers.Id = c.Id
//			answers.Description = c.Description
//			err := mongo.UpdateStoreAnswers(answers)
//			if err != nil {
//				log.Println(err)
//			}
//		}
//	}
//	close(done)
//}

//func UpdateRestaurantsAnswers(done chan struct{}) {
//	candidates := GetCandidates(limit, offset)
//	for _, c := range candidates {
//		var answers model.CandidateRestaurant
//		if c.Type == "restaurant" {
//			answers.Id = c.Id
//			answers.Description = c.Description
//			err := mongo.UpdateRestaurantsAnswers(answers)
//			if err != nil {
//				log.Println(err)
//			}
//		}
//	}
//	close(done)
//}

func UpdateDriversAnswers(done chan struct{}) {
	candidates := GetCandidates(limit, offset)
	for _, c := range candidates {
		var answers model.CandidateDriver
		if c.Type == "drivers" {
			answers.Id = c.Id
			answers.Description = c.Description
			err := mongo.UpdatesDriversAnswers(answers)
			if err != nil {
				log.Println(err)
			}
		}
	}
	close(done)
}

func GetRestaurantsAnswers(done chan struct{}) {

	candidates := GetCandidates(10000, 90000)

	for _, c := range candidates {

		var answers model.CandidateRestaurant

		if c.Type == "restaurant" {

			answers.Q1 = c.Position
			answers.Q2 = true

			answers.Id = c.Id
			answers.CandidateID = c.CandidateID

			//c.Industry = strings.Trim(c.Industry, "{}")
			c.Companies = strings.Trim(c.Companies, "{}")
			//c.Systems = strings.Trim(c.Systems, "{}")
			//c.ExperiencePrograms = strings.Trim(c.Systems, "{}")
			//c.ExperienceTools = strings.Trim(c.ExperienceTools, "{}")
			//c.TypesWork = strings.Trim(c.TypesWork, "{}")

			//c.Industry = strings.Trim(c.Industry, "\"")

			//c.Systems = strings.Trim(c.Systems, "\"")
			//c.ExperiencePrograms = strings.Trim(c.ExperiencePrograms, "\"")
			//c.ExperienceTools = strings.Trim(c.ExperienceTools, "\"")
			//c.TypesWork = strings.Trim(c.TypesWork, "\"")

			//if c.TypesWork == "" && c.ExperienceTools == "" && c.Systems == "" && c.ExperiencePrograms == "" && c.Companies == "" && c.Industry == "" {
			//	continue
			//}

			//answers.Q4 = strings.Split(c.Industry, ",")

			temp, err := strconv.Atoi(c.Experience)
			if err != nil {
				log.Println(err)
			}
			if temp != 0 {
				answers.Q5 = true
			} else {
				answers.Q5 = false
			}

			var newV []rune
			answers.Q6 = strings.Split(c.Companies, ",")
			for i, v := range answers.Q6 {
				for _, v2 := range v {
					if v2 == '\\' || v2 == '"' {
						continue
					}
					newV = append(newV, v2)
				}
				answers.Q6[i] = string(newV)
				newV = nil
			}

			//answers.Q7 = strings.Split(c.ExperiencePrograms, ",")
			//answers.Q8 = strings.Split(c.ExperienceTools, ",")
			//answers.Q9 = strings.Split(c.TypesWork, ",")

			//answers.Q10 = c.FindWorkTime
			//answers.Q11 = c.Salary

			err = mongo.InsertRestaurantAnswers(answers)
			if err != nil {
				log.Println(err)
			}
		}
	}
	close(done)
}

func GetDriversAnswers() {

	candidates := GetCandidates(10000, 90000)

	for _, c := range candidates {

		var answers model.CandidateDriver

		if c.Type == "drivers" {

			answers.Q1 = c.Position
			answers.Q2 = false // У вас есть водительские права?

			answers.Id = c.Id
			answers.CandidateID = c.CandidateID

			//c.Industry = strings.Trim(c.Industry, "{}")
			c.Companies = strings.Trim(c.Companies, "{}")

			//c.Systems = strings.Trim(c.Systems, "{}")
			//c.ExperiencePrograms = strings.Trim(c.Systems, "{}")
			//c.ExperienceTools = strings.Trim(c.ExperienceTools, "{}")
			//c.TypesWork = strings.Trim(c.TypesWork, "{}")

			//c.Industry = strings.Trim(c.Industry, "\"")

			//c.Systems = strings.Trim(c.Systems, "\"")
			//c.ExperiencePrograms = strings.Trim(c.ExperiencePrograms, "\"")
			//c.ExperienceTools = strings.Trim(c.ExperienceTools, "\"")
			//c.TypesWork = strings.Trim(c.TypesWork, "\"")

			//if c.TypesWork == "" && c.ExperienceTools == "" && c.Systems == "" && c.ExperiencePrograms == "" && c.Companies == "" && c.Industry == "" {
			//	continue
			//}

			//answers.Q4 = strings.Split(c.Industry, ",")

			// опыт работы погру	зчиком
			temp, err := strconv.Atoi(c.Experience)
			if err != nil {
				log.Println(err)
			}
			if temp != 0 {
				answers.Q6 = true
				answers.Q7 = temp
			} else {
				answers.Q6 = false
			}

			var newV []rune
			// компании
			answers.Q9 = strings.Split(c.Companies, ",")
			for i, v := range answers.Q9 {
				for _, v2 := range v {
					if v2 == '\\' || v2 == '"' {
						continue
					}
					newV = append(newV, v2)
				}
				answers.Q9[i] = string(newV)
				newV = nil
			}

			//answers.Q7 = strings.Split(c.ExperiencePrograms, ",")
			//answers.Q8 = strings.Split(c.ExperienceTools, ",")
			//answers.Q9 = strings.Split(c.TypesWork, ",")

			//answers.Q10 = c.FindWorkTime
			//answers.Q11 = c.Salary

			err = mongo.InsertDriversAnswers(answers)
			if err != nil {
				log.Println(err)
			}
		}
	}
}
