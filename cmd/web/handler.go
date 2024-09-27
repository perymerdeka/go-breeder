package main

import (
	"net/http"
)

func (app *application) showHome(w http.ResponseWriter, r *http.Request) {
	app.render(w, "home.page.gohtml", nil)
}