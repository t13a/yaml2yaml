package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
	"os"
)

const (
	jsonPrefix = ""
	jsonIndent = "  "
)

func main() {
	jsonOutput := flag.Bool("json-output", false, "output in JSON format instead of YAML")
	flag.Parse()

	in, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		log.Fatalf("%v", err)
	}

	out, err := Format(in, *jsonOutput)
	if err != nil {
		log.Fatalf("%v", err)
	}

	fmt.Print(string(out))
}

func Format(in []byte, jsonOutput bool) ([]byte, error) {
	if jsonOutput {
		return FormatJSON(in)
	} else {
		return FormatYAML(in)
	}
}

func FormatJSON(in []byte) (out []byte, err error) {
	var m interface{}
	err = yaml.Unmarshal(in, &m)
	if err != nil {
		return
	}
	bytes, err := json.MarshalIndent(convertGenericMapToStringMap(m), jsonPrefix, jsonIndent)
	if err != nil {
		return
	}
	out = []byte(string(bytes) + "\n")
	return
}

func FormatYAML(in []byte) (out []byte, err error) {
	var m interface{}
	err = yaml.Unmarshal(in, &m)
	if err != nil {
		return
	}
	return yaml.Marshal(&m)
}

func convertGenericMapToStringMap(in interface{}) interface{} {
	switch x := in.(type) {
	case map[interface{}]interface{}:
		m2 := map[string]interface{}{}
		for k, v := range x {
			switch k2 := k.(type) {
			case string:
				m2[k2] = convertGenericMapToStringMap(v)
			default:
				m2[k.(string)] = convertGenericMapToStringMap(v)
			}
		}
		return m2
	case []interface{}:
		for i, v := range x {
			x[i] = convertGenericMapToStringMap(v)
		}
	}
	return in
}
