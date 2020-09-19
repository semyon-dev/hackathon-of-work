package model

import (
	"fmt"
	"strings"
)

type Candidate struct {
	Id           int64  `db:"id"` // postgres primary key
	Type         string `db:"type"`
	Experience   string `db:"experience"`
	CandidateID  string `db:"candidate_id"`
	Position     string `db:"position"`
	Organization string `db:"organization"`
	Description  string `db:"description"`
	StartDate    string `db:"start_date"`
	EndDate      string `db:"end_date"`
}

type NewVacation struct {
	Id              string `db:"id"`
	Type            string `db:"type"`
	Name            string `db:"name"`
	Area            string `db:"area"`
	City            string `db:"city"`
	CompanyID       string `db:"company_id"`
	Company         string `db:"company"`
	CompanyLink     string `db:"company_link"`
	PublicationDate string `db:"pub_date"`
	Salary          string `db:"salary_from"`
	SalaryCurrency  string `db:"salary_currency"`
	Employer        string `db:"emp_name"`
	ScheduleName    string `db:"schedule_name"`
	Experience      string `db:"exp_name"`
	KeySkills       string `db:"key_skills"`
	Specs           string `db:"specializations"`
	Description     string `db:"description"`
	Duties          string `db:"duties"`
	Demands         string `db:"demands"`
}

func CreateTypes() ([]string, []string, []string) {

	restaurant := []string{"официант", "ресторан", "бариста", "фаст-фуд", "повар"}
	drivers := []string{"водитель", "машинист"}
	store := []string{"склад", "кладовщик", "комплектовщик"}

	return restaurant, drivers, store
}

func (v *NewVacation) Split(correct, incorrect *int) {
	pieces := strings.Split(strings.ToLower(v.Description), "обязанности:")

	var duty, demand string

	if len(pieces) > 1 {
		temp := pieces[1]
		if strings.Contains(temp, "требования:") {
			i := strings.Index(temp, "требования:")
			duty = temp[:i]
		} else {
			fmt.Println(v.Description)
			*incorrect++
			return
		}

		temps := strings.Split(strings.ToLower(temp), "требования:")
		if len(temps) > 1 {
			t := temps[1]
			if strings.Contains(t, "условия:") {
				i := strings.Index(t, "условия:")
				demand = t[:i]
			} else {
				fmt.Println(v.Description)
				*incorrect++
				return
			}

			v.Duties = duty
			v.Demands = demand
			*correct++
		} else {
			fmt.Println("Duty error")
		}
	}
}

// ------------------- ML -----------------

type Res []Data

// data to answer
type Data struct {
	Title      string      `json:"title"`
	Paragraphs []Paragraph `json:"paragraphs"`
}

type Paragraph struct {
	Context string
	Qas     []Qas `json:"qas"`
}

type Qas struct {
	Answers  []Answer `json:"answers"`
	Question string   `json:"question"`
	Id       string   `json:"id"`
}

type Answer struct {
	AnswerStart uint64 `json:"answer_start"`
	Text        string `json:"text"`
}
