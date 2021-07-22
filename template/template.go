package template

import (
	"fmt"
	"html/template"
	"os"
	"path"

	"github.com/keezeden/pocket/utilities"
)


func GenerateEntry(data interface{}, outpath *string) {
	wd, err := os.Getwd()
	t, err := template.ParseFiles(path.Join(wd, "template/template.html"))

	utilities.Check(err)

	f, err := os.Create(*outpath)
	err = t.Execute(f, data)

	utilities.Check(err)
	
	fmt.Printf("Wrote Entry to %d Successfully", &outpath)
}