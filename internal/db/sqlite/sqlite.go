package sqlite

import (
	"database/sql"

	"github.com/uforco/sqlitebygo/internal/config"
	_ "github.com/mattn/go-sqlite3"
)

type Sqlite struct {
	Storage *sql.DB
}

// connect database
func New(cfg *config.Config) (*Sqlite, error) {
	db, err := sql.Open("sqlite3", cfg.StoragePath)
	if err != nil {
		return nil, err
	}
	return &Sqlite{
		Storage: db,
	}, nil
}

// query row
func (db *Sqlite) QueryRow(query string, args ...interface{}) *sql.Row {
	return db.Storage.QueryRow(query, args...)
}

// query
func (db *Sqlite) Query(query string, args ...interface{}) (*sql.Rows, error) {
	stmt, err := db.Storage.Prepare(query)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	result, err := stmt.Query(args...)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// insert, update, delete
func (db *Sqlite) Exec(query string, args ...interface{}) (sql.Result, error) {
	stmt, err := db.Storage.Prepare(query)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	result, err := stmt.Exec(args...)
	if err != nil {
		return nil, err
	}
	return result, nil
}
