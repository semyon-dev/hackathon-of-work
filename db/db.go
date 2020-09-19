package db

import (
	"database/sql"
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

	bks := make([]*Book, 0)
	for rows.Next() {
		bk := new(Book)
		err := rows.Scan(&bk.isbn, &bk.title, &bk.author, &bk.price)
		if err != nil {
			log.Fatal(err)
		}
		bks = append(bks, bk)
	}
}
