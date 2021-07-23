package main

import (
	"flag"
	"fmt"
	"math"
	"strings"

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

	id := utilities.VersionToId(version)
	entry, species := api.GetPokemonByEntry(id)	

	data := struct {
		ID int
		Name string
		Species string
		ImageURL string
		Height float64
		Weight float64
		Description string
	}{
		ID: id,
		Name: fmt.Sprint(entry["name"]),
		Species: strings.Replace(fmt.Sprint(species["genera"].([]interface{})[7].(map[string]interface{})["genus"]), " Pokémon", "", -1),
		ImageURL: fmt.Sprint(entry["sprites"].(map[string]interface{})["versions"].(map[string]interface{})["generation-i"].(map[string]interface{})["red-blue"].(map[string]interface{})["front_default"]),
		Height: math.Round((entry["height"].(float64) / 3.048) * 100)/100,
		Weight: math.Round((entry["weight"].(float64) / 4.536)),
	}

	template.GenerateEntry(data, outpath)
}