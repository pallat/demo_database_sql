package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	sqliteVersion()
	sqliteCreateTable()
	sqliteSelectAll()
}

func sqliteVersion() {
	db, err := sql.Open("sqlite3", ":memory:")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	var version string
	err = db.QueryRow("SELECT SQLITE_VERSION()").Scan(&version)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(version)
}

func sqliteCreateTable() {
	db, err := sql.Open("sqlite3", "test.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	sts := `
DROP TABLE IF EXISTS todos;
CREATE TABLE todos(id INTEGER PRIMARY KEY, title TEXT, done BOOLEAN);
INSERT INTO todos(title, done) VALUES('Learn Go',true);
INSERT INTO todos(title, done) VALUES('Practice TDD',false);
INSERT INTO todos(title, done) VALUES('Make an API',false);
INSERT INTO todos(title, done) VALUES('Shop a book',false);
`
	_, err = db.Exec(sts)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("table todos created")
}

func sqliteSelectAll() {
	db, err := sql.Open("sqlite3", "test.db")

	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()

	rows, err := db.Query("SELECT * FROM todos")

	if err != nil {
		log.Fatal(err)
	}

	defer rows.Close()

	for rows.Next() {

		var id int
		var title string
		var done bool

		err = rows.Scan(&id, &title, &done)

		if err != nil {
			log.Fatal(err)
		}

		fmt.Printf("%d %s, status(%t)\n", id, title, done)
	}
}
