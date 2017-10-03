package main

import (
	"fmt"
	"os"
	"path/filepath"
	"text/template"
)

func process(result []ResultMap, structName string, packageName string) {

	otData := outputData{structName, packageName, result}
	cwd, err := os.Getwd()
	cmb := filepath.Join(cwd, "resources/type.tmpl")

	temp, err := template.ParseFiles(cmb)
	if err != nil {
		panic(err)
	}

	f, err := os.Create(fmt.Sprintf("%v.go", structName))
	defer f.Close()
	err = temp.ExecuteTemplate(f, "type", otData)
	if err != nil {
		panic(err)
	}

}

type outputData struct {
	StructName, PackageName string
	Columns                 []ResultMap
}
