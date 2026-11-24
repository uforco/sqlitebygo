package registar

import (
	"net/http"

	"github.com/uforco/sqlitebygo/internal/config"
	"github.com/uforco/sqlitebygo/internal/http/handlers/user"
)

func Api(router *http.ServeMux, clg *config.Config) {
	//user controler api registar
	user.UserControler(router)

}
