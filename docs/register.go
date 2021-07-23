package docs

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os/exec"
	"strings"
)

var routes []RouteInfo = []RouteInfo{}

type RouteInfo struct {
	Name  string
	Route string

	Description string
}

func AddDocs(route, funcName string) error {

	output, err := exec.Command("go", "doc", funcName).Output()

	if err != nil {
		return err
	}

	cleanDocs := []string{}

	// Remove whitespace at the start and end of each line
	for _, i := range strings.Split(string(output), "\n") {
		cleanDocs = append(cleanDocs, strings.TrimSpace(i))
	}

	// Remove empty lines. Output - A single line documentation (preferrably)
	documentation := strings.Trim(strings.Join(cleanDocs[3:], "\n"), "\n")

	routeDocs := RouteInfo{
		funcName,
		route,
		documentation,
	}

	if err != nil {
		log.Fatal(err)
	}

	routes = append(routes, routeDocs)

	return nil
}

func GenerateDocs() {

	jsonData := ""

	for _, route := range routes {
		docsJson, _ := json.MarshalIndent(route, "", "   ")
		jsonData += string(docsJson) + ",\n"
	}

	err := ioutil.WriteFile("../../docs/docs.json", []byte(jsonData), 0644)

	if err != nil {
		log.Fatal(err)
	}
}
