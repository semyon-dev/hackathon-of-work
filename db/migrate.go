package db

import (
	"fmt"
	"hackathon-work/model"
	"hackathon-work/transform"
)

func UpdateCandidates(candidates []model.Candidate) {

	restaurant, drivers, store := model.CreateTypes()

	c := 0
	i := 0
	for _, candidate := range candidates {
		isDetected := false
		candidate, isDetected = transform.DetectCandidateType(candidate, restaurant, "restaurant")
		if !isDetected {
			candidate, isDetected = transform.DetectCandidateType(candidate, drivers, "drivers")
			if !isDetected {
				candidate, isDetected = transform.DetectCandidateType(candidate, store, "store")
			}
		}
		if isDetected {
			c++
		} else {
			i++
		}
		UpdateCandidate(candidate)
	}

	fmt.Printf("\nCorrect candidate columns: %d\nIncorrect columns: %d\nIn percent: %0.2f%%\n", c, i, float64(c)/float64(100000)*100)

}

func UpdateVacations(vacations []model.NewVacation) {

	var correct, incorrect int

	restaurant, drivers, store := model.CreateTypes()

	for _, vacation := range vacations {
		vacation.Split(&correct, &incorrect)

		isDetected := false

		vacation, isDetected = transform.DetectVacationType(vacation, restaurant, "restaurant")
		if !isDetected {
			vacation, isDetected = transform.DetectVacationType(vacation, drivers, "drivers")
			if !isDetected {
				vacation, isDetected = transform.DetectVacationType(vacation, store, "store")
			}
		}
		UpdateNewVacation(vacation)
	}
	fmt.Printf("\nCorrect columns: %d\nIncorrect columns: %d\nIn percent: %0.2f%%\n", correct, incorrect, float64(correct)/float64(33000)*100)

}
