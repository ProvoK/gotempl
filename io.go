package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"strings"
	"text/template"
)

func includeFile(dir string) func(filePath string) (string, error) {
	return func(filePath string) (string, error) {
		data, err := ioutil.ReadFile(path.Join(dir, filePath))
		if err != nil {
			return "", err
		}
		return string(data), nil
	}
}

func readFromStdin() (string, error) {
	scanner := bufio.NewScanner(os.Stdin)
	var output strings.Builder

	for scanner.Scan() {
		output.WriteString(fmt.Sprintln(scanner.Text()))
	}

	if err := scanner.Err(); err != nil {
		return "", err
	}
	return output.String(), nil
}

func getTemplate(templatePath string) (*template.Template, error) {

	templateDir, templateFile := path.Split(templatePath)
	funcs := template.FuncMap{
		"include": includeFile(templateDir),
	}
	if templateFile != "" {
		t, err := template.New(templateFile).Funcs(funcs).ParseFiles(templatePath)
		return t, err
	}
	tmpl, err := readFromStdin()
	t, err := template.New("stdin").Funcs(funcs).Parse(tmpl)
	return t, err
}
