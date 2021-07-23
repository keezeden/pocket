package api

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/keezeden/pocket/src/utilities"
)


var entryEndpoint = "https://pokeapi.co/api/v2/pokemon"
var speciesEndpoint = "https://pokeapi.co/api/v2/pokemon-species"

func Request(url string) (response map[string]interface{}) {
	resp, err := http.Get(url)
	
	utilities.Check(err)

	body, err := ioutil.ReadAll(resp.Body)
	
	utilities.Check(err)

	data := make(map[string]interface{})
	json.Unmarshal(body, &data)
	
	utilities.Check(err)

	return data
}

func GetPokemonByEntry(entry int) (entryJson map[string]interface{}, speciesJson map[string]interface{}) {
	entryUrl := fmt.Sprintf("%s/%d", entryEndpoint, entry)
	speciesUrl := fmt.Sprintf("%s/%d", speciesEndpoint, entry)

	pokemonEntry := Request(entryUrl)
	pokemonSpecies := Request(speciesUrl)

	return pokemonEntry, pokemonSpecies
}