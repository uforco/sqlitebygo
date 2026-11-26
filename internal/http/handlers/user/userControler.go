package user

import (
	"net/http"

	"github.com/uforco/sqlitebygo/internal/dbConfig"
)

func UserControler(router *http.ServeMux, storage dbconfig.DB) {
	router.HandleFunc("POST /user", CreateUser(storage))
}
