package registar

import (
	"net/http"

	"github.com/uforco/sqlitebygo/internal/config"
	"github.com/uforco/sqlitebygo/internal/db"
	"github.com/uforco/sqlitebygo/internal/http/handlers/user"
)

func Api(router *http.ServeMux, clg *config.Config, storage db.DB) {
	//user controler api registar
	user.UserControler(router, storage)

}
