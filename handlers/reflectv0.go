package handlers

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
)

func ReflectV0(w http.ResponseWriter, r *http.Request) {
	writer := io.MultiWriter(w, os.Stderr)
	w.WriteHeader(200)

	writer.Write(b("Method:"))
	writer.Write(b(r.Method))

	writer.Write(b("\nHeaders:"))
	for k, vs := range r.Header {
		writer.Write(b(fmt.Sprintf("%s: %#v", k, vs)))
	}

	writer.Write(b("\nPath:"))
	writer.Write(b(r.URL.Path))

	writer.Write(b("\nQuery Params:"))
	for k, vs := range r.URL.Query() {
		writer.Write(b(fmt.Sprintf("%s: %#v", k, vs)))
	}

	writer.Write(b("\nBody:"))
	defer r.Body.Close()
	bodyContent, _ := ioutil.ReadAll(r.Body)
	writer.Write(bodyContent)
}
