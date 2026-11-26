package dbconfig

import "database/sql"

type DB interface {
	QueryRow(string, ...interface{}) *sql.Row
	Query(...interface{}) (*sql.Rows, error)
	Exec(...interface{}) (sql.Result, error)
	ReadSchema(string) DB
}
