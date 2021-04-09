package handlers

import (
	"html"
	"net/http"
	"strconv"
	"time"
)

const defaultRespCode = 200
const defaultRespBody = "OK"
const defaultRespDelay = "0s"

// The endpoint takes following query params:
//
// response_code int (200)
//
// response_body string ("OK")
//
// response_delay duration ("0s")
func ControlV0(w http.ResponseWriter, r *http.Request) {
	respCode := defaultRespCode
	respBody := defaultRespBody
	respDelay := defaultRespDelay

	for k, vs := range r.URL.Query() {
		switch k {
		case "response_code":
			respCode = force(strconv.Atoi(vs[0])).(int)
		case "response_body":
			respBody = vs[0]
		case "response_delay":
			respDelay = vs[0]
		}
	}

	time.Sleep(force(time.ParseDuration(respDelay)).(time.Duration))
	w.WriteHeader(respCode)
	w.Write(b(html.EscapeString(respBody)))
}
