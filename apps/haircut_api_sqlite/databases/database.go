package databases

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

var DB *sql.DB

func InitialDatabase() {
	var err error
	DB, err = sql.Open("sqlite3", "./haircut.db")
	if err != nil {
		log.Fatal(err)
	}

	createTableQuery := `
  CREATE TABLE IF NOT EXISTS haircuts(
  id INTEGER PRIMARY KEY AUTOINCREMENT,
  name TEXT NOT NULL,
  description TEXT,
  price REAL NOT NULL
  );`

	_, err = DB.Exec(createTableQuery)
	if err != nil {
		log.Fatal(err)
	}
}
