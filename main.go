package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"

	"gopkg.in/yaml.v3"
)

func main() {
	file, err := os.ReadFile("cv.yml")
	if err != nil {
		log.Fatal("failed to read cv file")
	}

	cv := CV{}
	err = yaml.Unmarshal(file, &cv)
	if err != nil {
		log.Fatal("failed to parse file")
	}
	s, err := json.MarshalIndent(cv, "", "\t")
	if err != nil {
		log.Fatal("failed to marshal json")
	}
	fmt.Printf("%s", s)
}
