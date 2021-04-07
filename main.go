package main

import (
	"fmt"
	"fredwangwang/testing-service/controllers"
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

	addr := fmt.Sprintf("%s:%s", host, port)
	fmt.Println("service listening on: http://" + addr)

	// https://github.com/gorilla/mux
	r := mux.NewRouter()
	r.HandleFunc("/hc", controllers.HcController)
	r.HandleFunc("/reflect", controllers.ReflectController)
	r.HandleFunc("/", controllers.RootController)

	srv := &http.Server{
		Handler: r,
		Addr:    addr,
		// Good practice: enforce timeouts for servers you create!
		WriteTimeout: 30 * time.Second,
		ReadTimeout:  30 * time.Second,
	}

	log.Fatal(srv.ListenAndServe())
}
