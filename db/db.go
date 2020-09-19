package db

import (
	"database/sql"
	"github.com/semyon-dev/HackatonWork/model"
	"log"

	_ "github.com/lib/pq"
)

func Connect() {
	connStr := "user=semyon dbname=dima sslmode=disable password=pubavi09 port=5432 connect_timeout=8 host=35.228.59.145"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}

	_ = db

	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	rows, err := db.Query("SELECT * FROM candidates_exp")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	resumes := make([]*model.Resume, 0)
	for rows.Next() {
		resume := new(model.Resume)
		err := rows.Scan(&resume.CandidateID, &resume.Description, &resume.Position, &resume.Organization)
		if err != nil {
			log.Fatal(err)
		}
		resumes = append(resumes, resume)
	}
}
