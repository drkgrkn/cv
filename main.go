package main

import (
	"context"
	"log"
	"os"

	"github.com/drkgrkn/cv/cv"
	"github.com/drkgrkn/cv/templates"
	"gopkg.in/yaml.v3"
)

func main() {
	ymlFile, err := os.ReadFile("cv.yml")
	if err != nil {
		log.Fatal("failed to read cv file")
	}

	cv := cv.CV{}
	err = yaml.Unmarshal(ymlFile, &cv)
	if err != nil {
		log.Fatal("failed to parse file")
	}

	c := templates.CV(cv)

	cvFile, err := os.Create("cv.html")
	if err != nil {
		log.Fatal("failed to create html file")
	}

	c.Render(context.Background(), cvFile)
}
