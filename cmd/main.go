package main

import (
	"fmt"
	"log"
	"os"

	v1schema "github.com/YanyChoi/manifesto/pkg/schema/v1"
	v1template "github.com/YanyChoi/manifesto/pkg/template/v1"
	flag "github.com/YanyChoi/manifesto/pkg/flag"
)

func main() {
	fmt.Println(`
                       __  ___             __         
 .--------.---.-.-----|__.'  _.-----.-----|  |_.-----.
 |        |  _  |     |  |   _|  -__|__ --|   _|  _  |
 |__|__|__|___._|__|__|__|__| |_____|_____|____|_____|
 `)

	flags := flag.NewFlags()

	schema := v1schema.NewV1Schema(flags.Template)
	t := v1template.NewTemplate(schema)
	manifest, err := t.Hydrate()
	if err != nil {
		log.Fatal(err)
	}

	err = os.WriteFile(flags.Output, manifest, 0644)
	if err != nil {
		log.Fatal(err)
	}
}
