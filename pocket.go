package main

import (
	"flag"
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/keezeden/pocket/api"
	"github.com/keezeden/pocket/changelog"
)

func main() {
	filepath := flag.String("file", "./CHANGELOG.md", "The file to transform into a pokedex entry.")
	flag.Parse()
	
	version, err := changelog.GetVersion(filepath)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(version)

	segments := strings.Split(version, ".")

	major, err := strconv.Atoi(segments[0])
	minor, err := strconv.Atoi(segments[1])

	entry := major + minor

	pokemon := api.GetPokemonByEntry(entry)


	fmt.Println(pokemon["name"])
}