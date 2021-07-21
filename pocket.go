package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/keezeden/pocket/api"
	"github.com/keezeden/pocket/changelog"
	"github.com/keezeden/pocket/utilities"
)

func main() {
	filepath := flag.String("file", "./CHANGELOG.md", "The file to transform into a pokedex entry.")
	flag.Parse()
	
	version, err := changelog.GetVersion(filepath)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(version)

	entry := utilities.VersionToEntry(version)
	pokemon := api.GetPokemonByEntry(entry)

	fmt.Println(pokemon["name"])
}