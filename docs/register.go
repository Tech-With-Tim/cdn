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

// Extract comments to form docs
func AddDocs(route, funcName string) error {

	rawFunc := funcName

	funcName = strings.ReplaceAll(funcName, " ", "")
	output, err := exec.Command("go", "doc", funcName).Output()

	if err != nil {
		return err
	}

	cleanDocs := []string{}

	// Remove whitespace at the start and end of each line
	for _, i := range strings.Split(string(output), "\n") {
		cleanDocs = append(cleanDocs, strings.TrimSpace(i))
	}

	// Remove empty lines. Output - A multiline doc string
	documentation := strings.Trim(strings.Join(cleanDocs[3:], "\n"), "\n")
	documentation = strings.ReplaceAll(documentation, "\n\n", "<br>")
	documentation = strings.ReplaceAll(documentation, "\n", " ")
	documentation = strings.ReplaceAll(documentation, "  ", " ")

	routeDocs := RouteInfo{
		rawFunc,
		route,
		documentation,
	}

	if err != nil {
		return err
	}

	routes = append(routes, routeDocs)

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
