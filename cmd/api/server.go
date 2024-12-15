package main

import (
	"fmt"
	"net/http"
)

func (app *application) serve() error {
	srv := http.Server{
		Addr:    fmt.Sprintf(":%d", app.config.port),
		Handler: nil,
	}
	return srv.ListenAndServe()
}
