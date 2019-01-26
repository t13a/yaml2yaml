package main

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
	"os"
)

func main() {
	in, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		log.Fatalf("%v", err)
	}

	out, err := Format(in)
	if err != nil {
		log.Fatalf("%v", err)
	}

	fmt.Printf("%s", string(out))
}

func Format(in []byte) (out []byte, err error) {
	var m interface{}
	err = yaml.Unmarshal(in, &m)
	if err != nil {
		return
	}
	return yaml.Marshal(&m)
}
