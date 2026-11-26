package user

import (
	"encoding/json"
	"errors"
	"io"
	"log/slog"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/uforco/sqlitebygo/internal/dbConfig"
	"github.com/uforco/sqlitebygo/types"
	errorhandle "github.com/uforco/sqlitebygo/utils/error-handle"
	"github.com/uforco/sqlitebygo/utils/response"
)

func CreateUser(db dbconfig.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		var user types.UserSchema

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
		if err := validator.New().Struct(user); err != nil {
			validationError := err.(validator.ValidationErrors)
			response.WriteJson(w, http.StatusBadRequest, errorhandle.ValidationError(validationError))
			return
		}

		// read sql file
		

		// execute sql query

		slog.Info("create a new ", "user", user)
		response.WriteJson(w, http.StatusCreated, map[string]interface{}{"message": "user created"})
	}
}
