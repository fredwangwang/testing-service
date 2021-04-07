package controllers

import (
	"net/http"
)

func RootController(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(200)
	w.Write(b("if you see this page, at least something is working"))
}
