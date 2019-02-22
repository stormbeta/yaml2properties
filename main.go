package main

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strconv"
)

// Inject from build
var commitVersion string

// Takes result of yaml parse and returns a flat property map using dot-separated paths
// IMPORTANT: Requires that the yaml file consist solely of maps and primitives, no arrays
//            Go's json package can't validate that for us as Go has no concept of higher-level types
func flatten(propertiesTree map[string]interface{}) map[string]string {
	mapped := make(map[string]string)
	for k, v := range propertiesTree {
		switch v.(type) {
		case map[interface{}]interface{}:
			// TODO: Why do I have to construct a new map instead of just being able to do a type assertion here?
			// Golang is really, really dumb about this
			promiseItIsAString := make(map[string]interface{})
			for key, val := range v.(map[interface{}]interface{}) {
				promiseItIsAString[key.(string)] = val
			}
			for innerKey, innerValue := range flatten(promiseItIsAString) {
				mapped[k+"."+innerKey] = innerValue
			}
		case string:
			mapped[k] = v.(string)
		case int:
			mapped[k] = strconv.Itoa(v.(int))
		case bool:
			mapped[k] = strconv.FormatBool(v.(bool))
		}
	}
	return mapped
}

func usage() {
	fmt.Fprintln(os.Stderr, "Usage: "+os.Args[0]+" FILENAME")
	fmt.Fprintln(os.Stderr, "Build Commit: "+commitVersion)
	os.Exit(1)
}

func main() {
	if len(os.Args) < 2 {
		usage()
	}

	filename := os.Args[1]

	filePath, err := filepath.Abs(filename)
	yamlFile, err := ioutil.ReadFile(filePath)
	if err != nil {
		log.Fatal(err)
		usage()
	}

	var parsed map[string]interface{}
	err = yaml.Unmarshal(yamlFile, &parsed)
	if err != nil {
		panic(err)
	}

	properties := flatten(parsed)
	for prop, path := range properties {
		fmt.Println(prop + "=" + path)
	}
}
