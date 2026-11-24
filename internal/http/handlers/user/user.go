package user

import (
	"log/slog"
	"net/http"
)

func CreateUser() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		slog.Info("create a new user....")
		w.Write([]byte("Hello you are "))
	}
}
