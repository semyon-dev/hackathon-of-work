package db

import (
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
	db, err = sqlx.Connect("postgres", "user=dima dbname=dima sslmode=disable password=pubavi09 port=5432 connect_timeout=8 host=35.228.59.145")
	if err != nil {
		log.Fatalln(err)
	}

	err = db.Ping()
	if err != nil {
		log.Fatal("Ping: ", err)
	}
}

func GetCandidates(limit, offset uint64) (candidates []model.Candidate) {
	err := db.Select(&candidates, "SELECT * FROM candidates_exp ORDER BY id LIMIT $1 OFFSET $2", limit, offset)
	if err != nil {
		log.Println(err)
	}

	return candidates
}

func GetVacations(limit, offset uint64) (vacancies []model.NewVacation) {
	err := db.Select(&vacancies, "SELECT * FROM vacancies ORDER BY id LIMIT $1 OFFSET $2", limit, offset)
	if err != nil {
		log.Println(err)
	}
	return vacancies
}

func UpdateNewVacation(vacancie model.NewVacation) {
	_, err := db.NamedExec(`update vacancies SET duties = :duties, demands = :demands, type = :type WHERE id = :id`, vacancie)
	if err != nil {
		log.Println(err)
	}
}

func UpdateCandidate(candidate model.Candidate) {
	_, err := db.NamedExec(`update candidates_exp SET type = :type WHERE id = :id`, candidate)
	if err != nil {
		log.Println(err)
	}
}
