package main

import (
	"fmt"
	"hackathon-work/api"
	"hackathon-work/db"
	"hackathon-work/model"
	"strings"
	"unicode"
)

func main() {

	db.Connect()

	restaurant, drivers, store := CreateTypes()

	go api.RunAPI()

	vacations := db.GetVacations(1, 11)

	for _, vacation := range vacations {
		fmt.Println(vacation.Id)
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

		vacation.Split()

		vacation = DetectType(vacation, restaurant, "restaurant")
		vacation = DetectType(vacation, drivers, "drivers")
		vacation = DetectType(vacation, store, "store")

		db.UpdateNewVacation(vacation)

		fmt.Printf("Type: %s,\nDuity: %s,\nDemand: %s.", vacation.Type, vacation.Duties, vacation.Demands)
	}

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
	drivers := []string{"водитель"}
	store := []string{"склад", "кладовщик", "комплектовщик"}

	return restaurant, drivers, store
}

func DetectType(vacation model.NewVacation, words []string, wantedType string) model.NewVacation {
	for _, word := range words {
		if strings.Contains(strings.ToLower(vacation.Name), word) {
			vacation.Type = wantedType
			break
		}
	}
	return vacation
}
