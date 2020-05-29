package main

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db, err := sql.Open("mysql", "root:development@/learning_go")

	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()

	tx, _ := db.Begin()
	stmt, _ := tx.Prepare("INSERT INTO users (id, name) VALUES (?, ?)")

	stmt.Exec(2000, "Daniela")
	stmt.Exec(2001, "Maria")
	stmt.Exec(2002, "Lídia")

	_, err = stmt.Exec(2000, "Vitor") // Duplicate key

	if err != nil {
		tx.Rollback()
		log.Fatal(err)
	}

	tx.Commit()
}
