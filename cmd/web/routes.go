package main

import (
	"net/http"

	"github.com/gorilla/mux"
)

func (app *application) routes() http.Handler {
	r := mux.NewRouter()
	r.Use(Cors)
	r.Use(app.RequestLog)
	r.HandleFunc("/", app.home).Methods(http.MethodGet)
	r.HandleFunc("/quote", app.createQuote).Methods(http.MethodPost, http.MethodOptions)

	return r
}
