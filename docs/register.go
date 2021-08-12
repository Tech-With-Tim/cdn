package docs

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os/exec"
	"strings"

	"gopkg.in/yaml.v2"
)

var routes []RouteInfo = []RouteInfo{}

type Docs struct {
	Content RouteInfo `yaml:"Content"`
}

type RouteInfo struct {
	Response      string `yaml:"Response"`
	URLParameters string `yaml:"URL Parameters"`

	Name  string
	Route string

	RequestBody []struct {
		Name string `yaml:"name"`
		Type string `yaml:"type"`
		Desc string `yaml:"Desc"`
	} `yaml:"Request Body"`

	Description string `yaml:"Description"`
}

// Extract comments to form docs
func AddDocs(route, funcName string) error {

	rawFunc := funcName

	funcName = strings.ReplaceAll(funcName, " ", "")
	output, err := exec.Command("go", "doc", funcName).Output()

	if err != nil {
		return err
	}

	rawDocs := "Content:\n"

	for num, line := range strings.Split(string(output), "\n") {
		if num >= 3 {
			rawDocs += line + "\n"
		}
	}

	source := []byte(rawDocs)

	var config Docs
	err = yaml.Unmarshal(source, &config)
	if err != nil {
		panic(err)
	}

	config.Content.Name = rawFunc
	config.Content.Route = route

	routes = append(routes, config.Content)

	return nil
}

// Write docs to a json file
func GenerateDocs() {

	jsonData, _ := json.MarshalIndent(routes, "", "   ")

	err := ioutil.WriteFile("../../docs/docs-template/public/docs.json", []byte(jsonData), 0644)

	if err != nil {
		log.Fatal(err)
	}
}
