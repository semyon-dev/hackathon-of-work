package db

import (
	"fmt"
	"hackathon-work/model"
	"log"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

var db *sqlx.DB

func Connect() {

	var err error

	// this Pings the database trying to connect, panics on error
	// use sqlx.Open() for sql.Open() semantics
	db, err = sqlx.Connect("postgres", "user=dima dbname=hack sslmode=disable password=pubavi09 port=5432 connect_timeout=8 host=35.228.59.145")
	if err != nil {
		log.Fatalln(err)
	}

	err = db.Ping()
	if err != nil {
		log.Fatal("Ping: ", err)
	}
}

func GetResumes(limit, offset uint64) (resumes []model.Resume) {
	err := db.Select(&resumes, "SELECT * FROM candidates_exp LIMIT $1 OFFSET $2", limit, offset)
	if err != nil {
		log.Fatal(err)
	}
	return resumes
}

func GetVacations(limit, offset uint64) (vacancies []model.Vacation) {
	err := db.Select(&vacancies, "SELECT * FROM vacancies LIMIT $1 OFFSET $2", limit, offset)
	if err != nil {
		log.Fatal(err)
	}
	return vacancies
}

func InsertNewResume(resume model.Resume) {
	result, err := db.NamedExec("INSERT INTO candidates_exp_new ()VALUES(:id, :)", resume)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(result.RowsAffected())
}

func InsertNewVacation(vacancie model.Vacation) {
	result, err := db.NamedExec("INSERT INTO vacancies_new")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(result.RowsAffected())
}
