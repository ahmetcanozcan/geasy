package server

import (
	"fmt"
	"net/http"
	"os"
	"os/exec"

	"github.com/gorilla/mux"
	"github.com/phayes/freeport"
)

// Run :
func Run(file string) {
	json := readJSONFile(file)
	r := mux.NewRouter()
	modulemap, _ := runModules(json.Modules)
	for _, route := range json.Routes {
		methods := route.Methods
		if len(methods) == 0 {
			methods = []string{"GET"}
		}
		r.
			HandleFunc(route.Path, generateCMDResultHandler(route.Cmd, modulemap, route.Params, route.Headers)).
			Methods(methods...)
	}

	staticFilePath := json.StaticFilePath
	if staticFilePath != "" {
		r.PathPrefix("/").Handler(http.FileServer(http.Dir(staticFilePath)))
	}

	http.ListenAndServe(":3000", r)
}

func generateCMDResultHandler(commands []string, modules map[string]string, params []string, headers []ResponseHeader) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		for _, header := range headers {
			w.Header().Add(header.Key, header.Value)
		}
		pparams := parseParams(params, modules, r)
		res := commands
		res = append(commands, pparams...)
		fmt.Println(res)
		cmd := exec.Command(res[0], res[1:]...)
		cmd.Stdout = w
		cmd.Stderr = os.Stderr
		cmd.Run()
		cmd.Wait()
	}

}

func runModules(modules []Module) (map[string]string, error) {
	ports, err := freeport.GetFreePorts(len(modules))
	result := make(map[string]string)
	if err != nil {
		return nil, err
	}
	fmt.Println("Empty Ports:", ports)
	for i, port := range ports {
		portStr := fmt.Sprintf("%d", port)
		mod := modules[i]
		commands := append(mod.Cmd, portStr)
		cmd := exec.Command(commands[0], commands[1:]...)
		cmd.Stdout = os.Stdout
		go cmd.Run()
		result[mod.Name] = portStr
	}

	return result, nil
}
