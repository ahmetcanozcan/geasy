package server

import (
	"fmt"
	"net/http"
	"os"
	"os/exec"

	"github.com/gorilla/mux"
)

// Run :
func Run(file string) {
	json := readJSONFile(file)
	r := mux.NewRouter()
	for _, route := range json.Routes {
		r.HandleFunc(route.Path, generateCMDResultHandler(route.Cmd, route.Params))
	}

	staticFilePath := json.StaticFilePath
	if staticFilePath != "" {
		r.PathPrefix("/").Handler(http.FileServer(http.Dir(staticFilePath)))
	}

	http.ListenAndServe(":3000", r)
}

func generateCMDResultHandler(commands []string, params []string) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		params = parseParams(params, r)
		res := commands
		res = append(commands, params...)
		fmt.Println(res)
		cmd := exec.Command(res[0], res[1:]...)
		cmd.Stdout = w
		cmd.Stderr = os.Stderr
		cmd.Run()
		cmd.Wait()
	}

}
