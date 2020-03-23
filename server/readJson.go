package server

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

// Route :
type Route struct {
	Path   string   `json:"path"`
	Params []string `json:"params"`
	Cmd    []string `json:"cmd"`
}

// Config :
type Config struct {
	StaticFilePath string  `json:"publicFiles"`
	Routes         []Route `json:"routes"`
	NotFound       string  `json:"notFound"`
}

func readJSONFile(filePath string) *Config {
	file, err := ioutil.ReadFile(filePath)
	if err != nil {
		fmt.Println("Read file error.")
		fmt.Println(err)
	}

	data := Config{}

	err = json.Unmarshal([]byte(file), &data)

	if err != nil {
		fmt.Println("JSON parsing error")
		fmt.Println(err)
	}

	return &data
}
