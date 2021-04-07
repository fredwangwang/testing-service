package controllers

import (
	"fmt"
	"io/ioutil"
	"net/http"
)


func ReflectController(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(200)

	w.Write(b("Method:"))
	w.Write(b(r.Method))

	w.Write(b("\nHeaders:"))
	for k, vs := range r.Header {
		w.Write(b(fmt.Sprintf("%s: %#v", k, vs)))
	}

	w.Write(b("\nPath:"))
	w.Write(b(r.URL.Path))

	w.Write(b("\nQuery Params:"))
	for k, vs := range r.URL.Query() {
		w.Write(b(fmt.Sprintf("%s: %#v", k, vs)))
	}

	w.Write(b("\nBody:"))
	defer r.Body.Close()
	bodyContent, _ := ioutil.ReadAll(r.Body)
	w.Write(bodyContent)
}
