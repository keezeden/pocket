package main

import (
	"flag"
	"fmt"
	"log"

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
}