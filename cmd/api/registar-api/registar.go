package registar

import (
	"net/http"

	"github.com/uforco/sqlitebygo/internal/config"
	dbconfig "github.com/uforco/sqlitebygo/internal/dbConfig"
	"github.com/uforco/sqlitebygo/internal/http/handlers/user"
)

func Api(router *http.ServeMux, clg *config.Config, storage dbconfig.DB) {
	//user controler api registar
	user.UserControler(router, storage)

}
