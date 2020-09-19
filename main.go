package main

import (
	"fmt"
	"hackathon-work/api"
	"hackathon-work/db"
	"hackathon-work/model"
	"strings"
	"sync"
	"unicode"
)

func main() {

	db.Connect()

	restaurant, drivers, store := CreateTypes()

	go api.RunAPI()
	var correct, incorrect int
	var limit, limitR uint64
	limit = 1000
	limitR = 100
	vacations := db.GetVacations(limit, 11)
	candidates := db.GetResumes(limitR, 0)
	wg := new(sync.WaitGroup)
	wg.Add(1)
	for _, vacation := range vacations {
		//fmt.Println("++++++",vacation.Description, "=========")
		//for _, word := range duties {
		//	if strings.Contains(strings.ToLower(vacation.Description), word) {
		//		index := strings.Index(strings.ToLower(vacation.Description), word)
		//		fmt.Println(index)
		//		temp := []rune(vacation.Description)
		//		t := temp[index:]
		//		t = t[len([]rune(duties[0])):]
		//		t = DeleteStartSpaces(t)
		//		fmt.Println("Parsed Duties: ", string(t))
		//	}
		//}

		vacation.Split(&correct, &incorrect)

		isDetected := false

		vacation, isDetected = DetectVacationType(vacation, restaurant, "restaurant")
		if !isDetected {
			vacation, isDetected = DetectVacationType(vacation, drivers, "drivers")
			if !isDetected {
				vacation, isDetected = DetectVacationType(vacation, store, "store")
			}
		}

		db.UpdateNewVacation(vacation)
		//fmt.Printf("Type: %s,\nDuity: %s,\nDemand: %s.", vacation.Type, vacation.Duties, vacation.Demands)
	}
	fmt.Printf("\nCorrect columns: %d\nIncorrect columns: %d\nIn percent: %0.2f%%\n", correct, incorrect, float64(correct)/float64(limit)*100)
	wg.Done()
	wg.Add(1)
	go func() {
		c := 0
		i := 0
		for _, candidate := range candidates {
			isDetected := false
			candidate, isDetected = DetectCandidateType(candidate, restaurant, "restaurant")
			if !isDetected {
				candidate, isDetected = DetectCandidateType(candidate, drivers, "drivers")
				if !isDetected {
					candidate, isDetected = DetectCandidateType(candidate, store, "store")
				}
			}
			if isDetected {
				c++
			} else {
				i++
			}
		}
		fmt.Printf("\nCorrect candidate columns: %d\nIncorrect columns: %d\nIn percent: %0.2f%%\n", c, i, float64(c)/float64(limitR)*100)

		wg.Done()
	}()

	wg.Wait()

}

func DeleteStartSpaces(raw string) (formatted string) {
	c := 0
	for i := 0; !unicode.IsLetter(rune(raw[i])); i++ {
		c++
	}
	formatted = raw[c:]
	return formatted
}

func CreateTypes() ([]string, []string, []string) {

	restaurant := []string{"официант", "ресторан", "бариста", "фаст-фуд", "повар"}
	drivers := []string{"водитель", "машинист"}
	store := []string{"склад", "кладовщик", "комплектовщик"}

	return restaurant, drivers, store
}

func DetectVacationType(vacation model.NewVacation, words []string, wantedType string) (model.NewVacation, bool) {
	var isDetected bool
	for _, word := range words {
		if strings.Contains(strings.ToLower(vacation.Name), word) {
			vacation.Type = wantedType
			isDetected = true
			break
		}
	}
	return vacation, isDetected
}

func DetectCandidateType(candidate model.Candidate, words []string, wantedType string) (model.Candidate, bool) {
	var isDetected bool
	for _, word := range words {
		if strings.Contains(strings.ToLower(candidate.Position), word) {
			candidate.Type = wantedType
			isDetected = true
			break
		}
	}
	return candidate, isDetected
}
