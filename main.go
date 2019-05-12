package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
)

type templateVars map[string]string

var vars = make(templateVars)
var templateFile string
var outputFile string

func (tv templateVars) String() string {
	return fmt.Sprintf("%+v", map[string]string(tv))
}

func init() {
	flag.StringVar(&templateFile, "template", "", "path for template file")
	flag.StringVar(&templateFile, "t", "", "path for template file (shortand)")
	flag.StringVar(&outputFile, "out", "", "path for template file")
	flag.StringVar(&outputFile, "o", "", "path for template file (shortand)")
	flag.Var(&vars, "var", "blah")
}

func (tv templateVars) Set(value string) error {
	if strings.Contains(value, "=") {
		splitted := strings.Split(value, "=")
		tv[splitted[0]] = splitted[1]
		return nil
	}
	tv[value] = ""
	return nil
}

func main() {
	flag.Parse()

	t, err := getTemplate(templateFile)
	if err != nil {
		log.Fatalf("Error parsing template: %s", err)
	}

	writer := os.Stdout
	if outputFile != "" {
		writer, err = os.Create(outputFile)
		if err != nil {
			log.Fatalf("Failed to create file %s: %s", outputFile, err)
		}
	}
	defer writer.Close()
	if err = t.Execute(writer, vars); err != nil {
		log.Fatalf("Failed execute: %s", err)
	}
}
