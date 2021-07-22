package api

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/keezeden/pocket/utilities"
)


var endpoint = "https://pokeapi.co/api/v2/pokemon"

func GetPokemonByEntry(entry int) (entryJson map[string]interface{}) {
	request := fmt.Sprintf("%s/%d", endpoint, entry)
	resp, err := http.Get(request)
	
	utilities.Check(err)

	body, err := ioutil.ReadAll(resp.Body)
	
	utilities.Check(err)

	data := make(map[string]interface{})
	json.Unmarshal(body, &data)
	
	utilities.Check(err)

	return data
}