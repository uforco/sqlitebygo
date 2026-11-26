package inittables

import (
	"database/sql"
	"log"
	"log/slog"

	readsqlfile "github.com/uforco/sqlitebygo/utils/readSqlFile"
)

func InitTables(db *sql.DB) {
	initTable("./sql/user/createTable.sql", db)
}

func initTable(path string, db *sql.DB) {
	var count int64
	var err error = nil
	if err := db.Ping(); err != nil {
		slog.Error("failed to ping db", "error", err)
		log.Fatal(err)
	}

	query, err := readsqlfile.ReadSqlFile(path)
	if err != nil {
		log.Fatal(err)
	}
	result, err := db.Exec(query)
	if err != nil {
		log.Fatal(err)
	}
	count, err = result.LastInsertId()
	if err != nil {
		log.Fatal(err)
	}
	count, err = result.RowsAffected()
	if err != nil {
		log.Fatal(err)
	}
	if count > 0 {
		slog.Info("successfully created table", "table", path)
	} else {
		slog.Info("table already exists", "table", path)
	}
}
