package main

import (

	"html/template"
	"net/http"
	"strconv"
	"path/filepath"

)

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
	var homeVikings []Viking
	db.Preload("Country").Find(&homeVikings)

	tpl, err := template.ParseFiles(filepath.Join(templateDir, "index.html"))
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	tpl.Execute(w, homeVikings)
}


func homeHandler2(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	finalId, err:= strconv.Atoi(id)
	if err != nil {
		http.Error(w, "ID invalide", http.StatusBadRequest)
		return
	}

	var theviking Viking
	db.Preload("Country").Preload("Events").First(&theviking,finalId)

	tpl, err := template.ParseFiles(filepath.Join(templateDir, "viking.html"))
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	tpl.Execute(w, theviking)
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

// country page
func homeHandler4(w http.ResponseWriter, r *http.Request) {
	var countries []Country
	db.Preload("Viking").Find(&countries)

	tpl, err := template.ParseFiles(filepath.Join(templateDir, "pays.html"))
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	tpl.Execute(w, countries)
}

func homeHandler5(w http.ResponseWriter, r *http.Request) {
	var events []Event
	err := db.Preload("Vikings").Find(&events).Error
    if err != nil {
        http.Error(w, "Erreur lors de la récupération des données", 500)
        return
    }

    tpl, err := template.ParseFiles(filepath.Join(templateDir, "event.html"))
    if err != nil {
        http.Error(w, err.Error(), 500)
        return
    }
	tpl.Execute(w, events)
}

func homeHandler6(w http.ResponseWriter, r *http.Request) {

	if r.Method == http.MethodPost {
		switch r.FormValue("action") {
		case "create":
			addTable(r.FormValue("name"),
					r.FormValue("image"),
					r.FormValue("burthyear"),
					r.FormValue("deadyear"),
					r.FormValue("periode"),
					r.FormValue("country_id"),
					r.FormValue("mdp"))
		case "update":
			updateTable(r.FormValue("viking_id1"),
					r.FormValue("name1"),
					r.FormValue("image1"),
					r.FormValue("burthyear1"),
					r.FormValue("deadyear1"),
					r.FormValue("periode1"),
					r.FormValue("country_id1"),
					r.FormValue("mdp"))
		case "delete":
			destroyTable(r.FormValue("name2"))
		}

		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}



	tpl, err := template.ParseFiles(filepath.Join(templateDir, "update.html"))
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	tpl.Execute(w, nil)
}