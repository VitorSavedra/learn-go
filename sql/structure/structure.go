package main

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

func exec(db *sql.DB, sql string) sql.Result {
	result, err := db.Exec(sql)

	if err != nil {
		panic(err)
	}

	return result
}

func main() {
	db, err := sql.Open("mysql", "root:development@/")

	if err != nil {
		panic(err)
	}

	defer db.Close()

	exec(db, "CREATE DATABASE IF NOT EXISTS learning_go")
	exec(db, "USE learning_go")
	exec(db, "DROP TABLE IF EXISTS users")
	exec(db, `CREATE TABLE users (
			id INTEGER AUTO_INCREMENT,
			name VARCHAR(80),
			PRIMARY KEY (id)
		)`)
}
