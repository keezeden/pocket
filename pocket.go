package main

import (
	"flag"
	"fmt"

	"github.com/keezeden/pocket/src/api"
	"github.com/keezeden/pocket/src/changelog"
	"github.com/keezeden/pocket/src/template"
	"github.com/keezeden/pocket/src/utilities"
)

func main() {
	inpath := flag.String("in", "./CHANGELOG.md", "The file to transform into a pokedex entry.")
	outpath := flag.String("out", "./pokedex.html", "The file to output the pokedex entry to.")
	flag.Parse()
	
	version, err := changelog.GetVersion(inpath)

	utilities.Check(err)

	fmt.Println(version)

	entry := utilities.VersionToEntry(version)
	pokemon := api.GetPokemonByEntry(entry)

	data := struct {
		Name string
	}{
		Name: fmt.Sprint(pokemon["name"]),
	}

	template.GenerateEntry(data, outpath)
}