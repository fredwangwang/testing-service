package handlers

import (
	"net/http"
)

func Root(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(200)
	w.Write(b("if you see this page, at least something is working"))
}
