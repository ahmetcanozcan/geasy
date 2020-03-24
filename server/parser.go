package server

import (
	"net/http"
	"strings"

	"github.com/gorilla/mux"
)

func parseParams(params []string, modules map[string]string, r *http.Request) []string {
	urlParams := mux.Vars(r)
	parseError := r.ParseForm()

	result := make([]string, len(params))
	for i, param := range params {
		// if param string starts with url:
		if strings.HasPrefix(param, "url:") {
			qu := param[4:]
			result[i] = urlParams[qu]
		} else if strings.HasPrefix(param, "body:") && parseError == nil {
			qu := param[5:]
			result[i] = r.FormValue(qu)
		} else if strings.HasPrefix(param, "mod:") {
			qu := param[4:]
			result[i] = modules[qu]
		} else {
			result[i] = params[i]
		}
	}
	return result
}
