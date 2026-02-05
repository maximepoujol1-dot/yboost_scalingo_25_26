package api

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

//get the API of all artist

func ChooseArtisteAll() ([]byte, error) {

	url := "https://apiyboostviking.osc-fr1.scalingo.io/data/viking.json"
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

//get the API of just one artist

func ChooseArtistePrecise(id string) ([]byte, error) {

	url := "https://apiyboostviking.osc-fr1.scalingo.io/data/viking.json/" + id
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
