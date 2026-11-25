package user

import "net/http"

func UserControler(router *http.ServeMux) {
	router.HandleFunc("POST /user", CreateUser())
}
