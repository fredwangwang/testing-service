package main

import (
	"fmt"
	"fredwangwang/testing-service/handlers"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gorilla/mux"
)

func getEnvOrDefault(k string, dv string) string {
	v := os.Getenv(k)
	if v == "" {
		return dv
	}
	return v
}

func main() {
	host := getEnvOrDefault("HOST", "0.0.0.0")
	port := getEnvOrDefault("PORT", "8000")
	pathPrefix := getEnvOrDefault("PATH_PREFIX", "/")

	addr := fmt.Sprintf("%s:%s", host, port)
	fmt.Println("service listening on: http://" + addr + pathPrefix)

	// https://github.com/gorilla/mux
	r := mux.NewRouter()

	s := r.PathPrefix(pathPrefix).Subrouter()

	s.HandleFunc("/v0/hc", handlers.HcV0)
	s.PathPrefix("/v0/reflect").HandlerFunc(handlers.ReflectV0)
	s.HandleFunc("/v0/control", handlers.ControlV0)

	s.HandleFunc("/", handlers.Root)
	s.HandleFunc("", handlers.Root)

	srv := &http.Server{
		Handler: r,
		Addr:    addr,
		// Good practice: enforce timeouts for servers you create!
		WriteTimeout: 30 * time.Second,
		ReadTimeout:  30 * time.Second,
	}

	log.Fatal(srv.ListenAndServe())
}
