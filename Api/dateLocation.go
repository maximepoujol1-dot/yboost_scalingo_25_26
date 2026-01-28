package api

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"encoding/json"
)



type LocationConcert struct {
	Locations2 []string `json:"locations"`
}



type DateConcert struct {
	Dates []string `json:"dates"`
}

//get the APi of Date and Location

func ChooseDateLocation(lien string) ([]byte, error) {

	url := lien
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



func RecupDate(lien string) []string {
	var D DateConcert
	body, err := ChooseDateLocation(lien)
	err = json.Unmarshal(body, &D)
	if err != nil {
		fmt.Println("Error unmarshalling json:", err)
	}
	return D.Dates
}



func RecupLocation(lien string) []string {
	var L LocationConcert
	body, err := ChooseDateLocation(lien)
	err = json.Unmarshal(body, &L)
	if err != nil {
		fmt.Println("Error unmarshalling json:", err)
	}
	return L.Locations2
}
