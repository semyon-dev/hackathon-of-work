package model

type Resume struct {
	CandidateID  string `db:"candidate_id"`
	Position     string `db:"position"`
	Organization string `db:"organization"`
	Description  string `db:"description"`
	StartDate    string `db:"start_date"`
	EndDate      string `db:"end_date"`
}

type NewResume struct {
	CandidateID  string `db:"candidate_id"`
	Position     string `db:"position"`
	Organization string `db:"organization"`
	Description  string `db:"description"`
	StartDate    string `db:"start_date"`
	EndDate      string `db:"end_date"`
}

type Vacation struct {
	Id              string `db:"id"`
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
}

type NewVacation struct {
	Id              string `db:"id"`
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
