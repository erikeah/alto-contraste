package main

import (
	"encoding/json"
	"io"
	"os"
	"text/template"
)

func main() {
	if len(os.Args) != 2 {
		os.Exit(1)
	}
	theme, err := os.Open(os.Args[1])
	if err != nil {
		println(err.Error())
		os.Exit(1)
	}
	decoder := json.NewDecoder(theme) // TODO: Read from file provided from param
	palette := &Palette{}
	if err := decoder.Decode(palette); err != nil {
		println(err.Error())
		os.Exit(1)
	}
	funcMap := template.FuncMap{}
	templateBytes, err := io.ReadAll(os.Stdin)
	if err != nil {
		println(err.Error())
		os.Exit(2)

	}
	templateContent := string(templateBytes)
	tmpl, err := template.New("").Funcs(funcMap).Parse(templateContent)
	if err != nil {
		println(err.Error())
		os.Exit(2)
	}
	err = tmpl.Execute(os.Stdout, palette)
	if err != nil {
		println(err.Error())
		os.Exit(2)
	}
}
