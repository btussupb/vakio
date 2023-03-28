package db

import (
	"database/sql"
	"fmt"

	"github.com/btussupb/vakio/internal/config"
	_ "github.com/mattn/go-sqlite3"
)

func OpenDb(cfg config.Database) (*sql.DB, error) {
	db, err := sql.Open("sqlite3", cfg.DBname)
	if err != nil {
		fmt.Println(err)
	}
	_, err = db.Exec("CREATE TABLE IF NOT EXISTS users (id INTEGER PRIMARY KEY AUTOINCREMENT, city TEXT, name TEXT, number TEXT, createdTime DATE)")
	if err != nil {
		fmt.Println(err)
	}
	// statement.Exec()
	return db, nil
}
