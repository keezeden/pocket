package template

import (
	"bytes"
	"fmt"
	"html/template"
	"log"
	"os"
	"path"
)


func GenerateEntry(data interface{}) {
	wd, err := os.Getwd()
	t, err := template.ParseFiles(path.Join(wd, "template/template.html"))

	if err != nil {
		log.Fatal(err)
	}

	var tpl bytes.Buffer
	errr := t.Execute(&tpl, data)

	if errr != nil {
		log.Fatal(errr)
	}

	fmt.Print(tpl.String())
}