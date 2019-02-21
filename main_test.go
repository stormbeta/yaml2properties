package main

import (
	"fmt"
	yaml2 "gopkg.in/yaml.v2"
	"path/filepath"
	"reflect"
	"runtime"
	"testing"
)

// Credit: https://github.com/benbjohnson/testing/blob/master/README.md
func equals(tb testing.TB, exp, act interface{}) {
	if !reflect.DeepEqual(exp, act) {
		_, file, line, _ := runtime.Caller(1)
		fmt.Printf("\033[31m%s:%d:\n\n\texp: %#v\n\n\tgot: %#v\033[39m\n\n", filepath.Base(file), line, exp, act)
		tb.FailNow()
	}
}

func TestConvertYamlToProperties(t *testing.T) {
	yamlText := `
a:
  b: leaf2
  c:
    d: leaf3
e: leaf1
types:
  integer: 3
  boolean: true
`
	yaml := make(map[string]interface{})
	err := yaml2.Unmarshal([]byte(yamlText), &yaml)
	if err != nil {
		panic(err)
	}
	propertiesMap := mapProperties(yaml)
	expected := map[string]string{
		"a.b":           "leaf2",
		"a.c.d":         "leaf3",
		"e":             "leaf1",
		"types.integer": "3",
		"types.boolean": "true",
	}

	equals(t, expected, propertiesMap)
}
