package main

import (
	"flag"
	"fmt"

	"github.com/keezeden/pocket/api"
	"github.com/keezeden/pocket/changelog"
	"github.com/keezeden/pocket/template"
	"github.com/keezeden/pocket/utilities"
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