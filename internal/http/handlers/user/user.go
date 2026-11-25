package user

import (
	"encoding/json"
	"errors"
	"io"
	"log/slog"
	"net/http"

	"github.com/uforco/sqlitebygo/schema"
	"github.com/uforco/sqlitebygo/utils/response"
)

func CreateUser() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		var user schema.UserSchema

		// decode request body into user struct
		// if request body is empty or invalid json
		// return error
		err := json.NewDecoder(r.Body).Decode(&user)
		if errors.Is(err, io.EOF) {
			response.WriteJson(w, http.StatusBadRequest, "request body is empty")
			return
		}
		if err != nil {
			response.WriteJson(w, http.StatusInternalServerError, err.Error())
			return
		}
		// request validation

		slog.Info("create a new ", "user", user)
		response.WriteJson(w, http.StatusCreated, map[string]interface{}{"message": "user created"})
	}
}
