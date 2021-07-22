package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/keezeden/pocket/api"
	"github.com/keezeden/pocket/changelog"
	"github.com/keezeden/pocket/template"
	"github.com/keezeden/pocket/utilities"
)

func main() {
	inpath := flag.String("in", "./CHANGELOG.md", "The file to transform into a pokedex entry.")
	outpath := flag.String("in", "./CHANGELOG.md", "The file to transform into a pokedex entry.")
	flag.Parse()
	
	version, err := changelog.GetVersion(inpath)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(version)

	entry := utilities.VersionToEntry(version)
	pokemon := api.GetPokemonByEntry(entry)

	name := fmt.Sprint(pokemon["name"])

	data := struct {
		Name string
	}{
		Name: name,
	}

	template.GenerateEntry(data, outpath)
}