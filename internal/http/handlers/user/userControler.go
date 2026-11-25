package user

import (
	"net/http"

	"github.com/uforco/sqlitebygo/internal/db"
)

func UserControler(router *http.ServeMux, storage db.DB) {
	router.HandleFunc("POST /user", CreateUser(storage))
}
