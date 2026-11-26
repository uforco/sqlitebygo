package sqlite

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
	"github.com/uforco/sqlitebygo/internal/config"
	readsqlfile "github.com/uforco/sqlitebygo/utils/readSqlFile"
)

type Sqlite struct {
	Storage *sql.DB
	content string
	err     error
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
func (s *Sqlite) Query(args ...interface{}) (*sql.Rows, error) {

	if s.err != nil {
		return nil, s.err
	}

	stmt, err := s.Storage.Prepare(s.content)
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
func (s *Sqlite) Exec(args ...interface{}) (sql.Result, error) {
	// get sql from file if any error check
	if s.err != nil {
		return nil, s.err
	}

	stmt, err := s.Storage.Prepare(s.content)
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

// read sql file
func (db *Sqlite) ReadSchema(path string) *Sqlite {
	contain, err := readsqlfile.ReadSqlFile(path)

	return &Sqlite{
		Storage: db.Storage,
		content: contain,
		err:     err,
	}

}
