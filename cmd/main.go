package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	v1schema "github.com/YanyChoi/manifesto/pkg/schema/v1"
	v1template "github.com/YanyChoi/manifesto/pkg/template/v1"
)

func main() {
	fmt.Println(`
                       __  ___             __         
 .--------.---.-.-----|__.'  _.-----.-----|  |_.-----.
 |        |  _  |     |  |   _|  -__|__ --|   _|  _  |
 |__|__|__|___._|__|__|__|__| |_____|_____|____|_____|
 `)

	var template string
	flag.StringVar(&template, "template", "template.yaml", "path to the template file")
	var output string
	flag.StringVar(&output, "output", "deployment.yaml", "path to the output file")
	flag.Parse()

	schema := v1schema.NewV1Schema(template)
	t := v1template.NewTemplate(schema)
	manifest, err := t.Hydrate()
	if err != nil {
		log.Fatal(err)
	}

	err = os.WriteFile(output, manifest, 0644)
	if err != nil {
		log.Fatal(err)
	}
}
