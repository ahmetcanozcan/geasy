package server

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

// ResponseHeader :
type ResponseHeader struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

// Route :
type Route struct {
	Methods []string         `json:"methods"`
	Headers []ResponseHeader `json:"headers"`
	Path    string           `json:"path"`
	Params  []string         `json:"params"`
	Cmd     []string         `json:"cmd"`
}

// Module :
type Module struct {
	Name string   `json:"name"`
	Cmd  []string `json:"cmd"`
}

// Option :
type Option struct {
	UseLogger bool `json:"useLogger"`
}

// Config :
type Config struct {
	Options        Option   `json:"options"`
	StaticFilePath string   `json:"publicDir"`
	Routes         []Route  `json:"routes"`
	NotFound       string   `json:"notFound"`
	Modules        []Module `json:"modules"`
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
