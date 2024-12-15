package main

import (
	"log/slog"
	"net/http"

	"github.com/vihaan404/metaverse-go/internal/database"
	"golang.org/x/crypto/bcrypt"
)

func (app *application) signinHandler(w http.ResponseWriter, r *http.Request) {
	return
}

func (app *application) signupHandler(w http.ResponseWriter, r *http.Request) {
	input := database.CreateUserParams{}

	err := app.readJSON(w, r, &input)
	if err != nil {
		slog.Error(err.Error())
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	password := input.Password

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		slog.Error(err.Error())
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	user, err := app.config.query.CreateUser(r.Context(), input)

	return
}
