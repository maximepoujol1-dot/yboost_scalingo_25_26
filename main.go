package main

import (
	"encoding/json"
	"fmt"
	api "viking-tracker/Api"
	"html/template"
	"net/http"
	"os"
	"path/filepath"
)


// struct for the API
type Person struct {
    ID            int    `json:"id"`
    Image         string `json:"image"`
    Name          string `json:"name"`
    BirthLocation string `json:"birthLocation"`
    Time          string `json:"time"`
    Birthday      int    `json:"birthday"`
    Deadyears     int    `json:"deadyears"`
}


var id string
var templateDir string

// truth table of filter
var verif = map[string]bool{
        "verifDanemark":    false,
		"verifSuède":    false,
		"verifNorvège":    false,
        "verifBirthday":         false,
        "verifDeadyears": false,
		"verifDébut": false,
		"verifExpansion": false,
		"verifApoge": false,
		"verifFin": false,


		"réinitialiser":     false,
    }


func main() {
	staticDir := "./front"
	templateDir = "./front/page"
	if exePath, err := os.Executable(); err == nil {
		candidate := filepath.Join(filepath.Dir(exePath), "front")
		if _, statErr := os.Stat(candidate); statErr == nil {
			staticDir = candidate
			templateDir = filepath.Join(candidate, "page")
		}
	}

	fsStatic := http.FileServer(http.Dir(staticDir))
	http.Handle("/static/", http.StripPrefix("/static/", fsStatic))

	http.HandleFunc("/redirect", redirectHandler)
	http.HandleFunc("/", homeHandler)
	http.HandleFunc("/viking", homeHandler2)
	http.HandleFunc("/aboutus", homeHandler3)
	http.HandleFunc("/filtre", homeHandler4)


	port := os.Getenv("PORT")
	if port == "" {
		port = "8081" // pour ton local seulement
	}

	fmt.Println("Server started at :" + port)

	err := http.ListenAndServe(":"+port, nil)
	if err != nil {
		fmt.Println("Server error:", err)
	}
}



// redirect to the good page
func redirectHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		page := r.FormValue("page")
		if page != "" {
			http.Redirect(w, r, page, http.StatusSeeOther)
			return
		}
	}
	http.Redirect(w, r, "/", http.StatusSeeOther)
}

// the principale page
func homeHandler(w http.ResponseWriter, r *http.Request) {
	body, err := api.ChooseArtisteAll()
	var users []Person
	marshall1(body, &users)

	tpl, err := template.ParseFiles(filepath.Join(templateDir, "index.html"))
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	tpl.Execute(w, users)
}

// the atist page
func homeHandler2(w http.ResponseWriter, r *http.Request) {
	id = r.URL.Query().Get("id")
	body, err := api.ChooseArtistePrecise(id)
	var user Person
	marshall2(body, &user)

	tpl, err := template.ParseFiles(filepath.Join(templateDir, "viking.html"))
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	tpl.Execute(w, user)
}



// about us page
func homeHandler3(w http.ResponseWriter, r *http.Request) {

	tpl, err := template.ParseFiles(filepath.Join(templateDir, "aboutus.html"))
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	tpl.Execute(w, nil)
}

// filter page
func homeHandler4(w http.ResponseWriter, r *http.Request) {
	
	var user []Person
	var userSave []Person

	body, err := api.ChooseArtisteAll()
	marshall1(body, &user)
	userSave = user

	filter := r.URL.Query().Get("filter")
	if filter != "" {
		api.Verifswitch(verif, filter)
	}

	if verif["verifSuède"] == true{
		verif["verifNorvège"] = false
		verif["verifDanemark"] = false
		var results []Person
		for _, p := range user {
			if p.BirthLocation == "Suède" {
				results = append(results, p)
			}
		}
		user = results
	}

	if verif["verifNorvège"] == true{
		verif["verifSuède"] = false
		verif["verifDanemark"] = false
		var results []Person
		for _, p := range user {
			if p.BirthLocation == "Norvège" {
				results = append(results, p)
			}
		}
		user = results
	}

	if verif["verifDanemark"] == true{
		verif["verifSuède"] = false
		verif["verifNorvège"] = false
		var results []Person
		for _, p := range user {
			if p.BirthLocation == "Danemark" {
				results = append(results, p)
			}
		}
		user = results
	}

	if verif["verifDébut"] == true{
		verif["verifFin"] = false
		verif["verifExpansion"] = false
		verif["verifApoge"] = false
		var results []Person
		for _, p := range user {
			if p.Time == "Début de l'age des vikings" {
				results = append(results, p)
			}
		}
		user = results
	}

	if verif["verifExpansion"] == true{
		verif["verifApoge"] = false
		verif["verifFin"] = false
		verif["verifDébut"] = false
		var results []Person
		for _, p := range user {
			if p.Time == "Expansion des vikings" {
				results = append(results, p)
			}
		}
		user = results
	}

	if verif["verifApoge"] == true{
		verif["verifExpansion"] = false
		verif["verifFin"] = false
		verif["verifDébut"] = false
		var results []Person
		for _, p := range user {
			if p.Time == "Apogée des vikings" {
				results = append(results, p)
			}
		}
		user = results
	}

	if verif["verifFin"] == true{
		verif["verifApoge"] = false
		verif["verifExpansion"] = false
		verif["verifDébut"] = false
		var results []Person
		for _, p := range user {
			if p.Time == "Fin de l'age des vikings" {
				results = append(results, p)
			}
		}
		user = results
	}

	if verif["réinitialiser"] == true {
		user = userSave
		verif["verifDébut"] = false
		verif["verifExpansion"] = false
		verif["verifApoge"] = false
		verif["verifFin"] = false
		verif["verifNorvège"] = false
		verif["verifSuède"] =false		
		verif["verifDanemark"] = false

		verif["réinitialiser"] = false

	}

	tpl, err := template.ParseFiles(filepath.Join(templateDir, "filter.html"))
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	tpl.Execute(w, user)
}

// unmarshalls functions to create the struct

func marshall1(jsonFromApi []byte, truc *[]Person) {

	err := json.Unmarshal(jsonFromApi, truc)
	if err != nil {
		fmt.Println("Error unmarshalling json:", err)
	}
}


func marshall2(jsonFromApi []byte, truc *Person) {

	err := json.Unmarshal(jsonFromApi, truc)
	if err != nil {
		fmt.Println("Error unmarshalling json:", err)
	}
}