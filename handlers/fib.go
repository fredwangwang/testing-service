package handlers

import (
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// Fib the freaking slow way, good for burning some cycles
//
// Running on Ryzen 3600
//
// 45 -> 6.5s
//
// 40 -> 0.6s
//
// 36 -> 0.1s
func fib(n int) int {
	if n <= 1 {
		return n
	}
	return fib(n-1) + fib(n-2)
}

func Fib(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	res := fib(force(strconv.Atoi(vars["fib"])).(int))
	w.WriteHeader(200)
	w.Write(b(strconv.Itoa(res)))
}
