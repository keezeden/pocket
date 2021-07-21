package api

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)


var endpoint = "https://pokeapi.co/api/v2/pokemon"

func GetPokemonByEntry(entry int) (entryJson map[string]interface{}) {
	request := fmt.Sprintf("%s/%d", endpoint, entry)
	resp, err := http.Get(request)
	
	if err != nil {
        log.Fatal(err)
    }

	body, err := ioutil.ReadAll(resp.Body)
	
	if err != nil {
	   log.Fatalln(err)
	}

	data := make(map[string]interface{})
	json.Unmarshal(body, &data)
	
	if err != nil {
		log.Fatalln(err)
	 }

	return data
}