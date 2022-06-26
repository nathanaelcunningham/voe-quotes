package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func (app *application) home(w http.ResponseWriter, r *http.Request) {
	latestQuotes, err := app.quotes.Latest()
	if err != nil {
		app.serverError(w, err)
		return
	}
	jsonQuotes, err := json.Marshal(latestQuotes)
	if err != nil {
		app.serverError(w, err)
		return
	}

	w.Write(jsonQuotes)
}

func (app *application) createQuote(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "create quote")
}
