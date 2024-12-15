package main

import (
	"log/slog"
	"net/http"
)

func (app *application) healthcheck(w http.ResponseWriter, r *http.Request) {
	data := map[string]string{
		"status": "available",
	}
	err := app.writeJSON(w, r, http.StatusOK, envelope{"health ": data}, nil)
	if err != nil {
		slog.Error(err.Error())
		http.Error(w, "Internal server error", http.StatusInternalServerError)
	}
}
