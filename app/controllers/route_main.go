package controllers

import (
	"net/http"
)

func top(w http.ResponseWriter, r *http.Request) {
	GeneratHTML(w, "Hello", "layout", "top")
}
