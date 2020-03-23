package server

import (
	"net/http"
	"strings"

	"github.com/gorilla/mux"
)

func parseParams(params []string, r *http.Request) []string {
	urlParams := mux.Vars(r)
	for i, param := range params {
		// if param string starts with url:
		if strings.HasPrefix(param, "url:") {
			qu := param[4:]
			params[i] = urlParams[qu]
		}
	}
	return params
}
