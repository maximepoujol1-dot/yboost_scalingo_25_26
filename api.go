package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"encoding/json"
)

//get the API of all artist

func FetchAll(url string) ([]byte, error) {

	
	response, err := http.Get(url)

	if err != nil {

		fmt.Println("Erreur lors de la requête HTTP :", err)

		return nil, err

	}

	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)

	if err != nil {

		fmt.Println("Erreur lors de la lecture de la réponse :", err)

		return nil, err

	}

	return body, nil
}

func FetchPrecise(info string, id string) ([]byte, error) {

	url := info+"/"+id
	response, err := http.Get(url)

	if err != nil {

		fmt.Println("Erreur lors de la requête HTTP :", err)

		return nil, err

	}

	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)

	if err != nil {

		fmt.Println("Erreur lors de la lecture de la réponse :", err)

		return nil, err

	}

	return body, nil
}


func marshall1(jsonFromApi []byte, truc *[]Viking) {

	err := json.Unmarshal(jsonFromApi, truc)
	if err != nil {
		fmt.Println("Error unmarshalling json:", err)
	}
}

func marshall2(jsonFromApi []byte, truc *Viking) {

	err := json.Unmarshal(jsonFromApi, truc)
	if err != nil {
		fmt.Println("Error unmarshalling json:", err)
	}
}

func marshallCountry(jsonFromApi []byte, truc *[]Country) {

	err := json.Unmarshal(jsonFromApi, truc)
	if err != nil {
		fmt.Println("Error unmarshalling json:", err)
	}
}