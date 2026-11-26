package dbconfig

import "database/sql"

type DB interface {
	QueryRow(string, ...interface{}) (*sql.Row, error)
	Query(string, ...interface{}) (*sql.Rows, error)
	Exec(string, ...interface{}) (sql.Result, error)
	ReadSchema(string) map[string]interface{}
}
