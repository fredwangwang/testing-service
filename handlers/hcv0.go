package handlers

import (
	"fmt"
	"net/http"
	"strconv"
)

var hcStatusCode = 200

func HcV0(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		w.Header().Set("testing-key", "testing-val")
		w.WriteHeader(hcStatusCode)
		w.Write(b(fmt.Sprintf("this is the hc response with status code: %d", hcStatusCode)))
	}

	if r.Method == http.MethodPost {
		var value = r.URL.Query()["status"]
		if len(value) == 0 {
			w.WriteHeader(500)
			w.Write(b("no status value provided"))
			return
		}

		var statusCode, err = strconv.Atoi(value[0])
		if err != nil {
			w.WriteHeader(500)
			w.Write(b("failed to parse status value provided"))
			return
		}

		hcStatusCode = statusCode
		w.WriteHeader(200)
		w.Write(b(fmt.Sprintf("hc status code set to: %d", hcStatusCode)))
	}
}
