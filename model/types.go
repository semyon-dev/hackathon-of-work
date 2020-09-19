package model

type Resume struct {
	CandidateID  string `db:"candidate_id"`
	Position     string `db:"position"`
	Organization string `db:"organization"`
	Description  string `db:"description"`
	StartDate    string `db:"start_date"`
	EndDate      string `db:"end_date"`
}

type Vacation struct {
	Id        string
	Name      string // Название
	AreaName  string
	City      string
	CompanyID string
	Company   string
	Duties    string // Обязанности
	Demands   string // Требования
	Terms            // Условия
}

type Terms struct {
	ScheduleType string // График работы
	Timetable    string // Время работы
	Salary       string // Зарплата
}
