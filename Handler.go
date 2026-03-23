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
	
	if vikings == nil {
		db.Preload("Country").Find(&vikings)
	}

	tpl, err := template.ParseFiles(filepath.Join(templateDir, "index.html"))
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	tpl.Execute(w, vikings)
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

// filter page
func homeHandler4(w http.ResponseWriter, r *http.Request) {

	if len(pays) == 0 {
		db.Preload("Viking").Find(&pays)
	}

	tpl, err := template.ParseFiles(filepath.Join(templateDir, "pays.html"))
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	tpl.Execute(w, pays)
}

func homeHandler5(w http.ResponseWriter, r *http.Request) {
    
    err := db.Preload("Vikings").Find(&event).Error
    if err != nil {
        http.Error(w, "Erreur lors de la récupération des données", 500)
        return
    }

    tpl, err := template.ParseFiles(filepath.Join(templateDir, "event.html"))
    if err != nil {
        http.Error(w, err.Error(), 500)
        return
    }
    tpl.Execute(w, event)
}

func homeHandler6(w http.ResponseWriter, r *http.Request) {

	if r.Method == http.MethodPost {
		createTable(r.FormValue("name"),
					r.FormValue("image"),
					r.FormValue("burthyear"),
					r.FormValue("deadyear"),
					r.FormValue("periode"),
					r.FormValue("country_id"),
					r.FormValue("mdp"))
	}

	if r.Method == http.MethodPost {
		updateTable(r.FormValue("viking_id1"),
					r.FormValue("name1"),
					r.FormValue("image1"),
					r.FormValue("burthyear1"),
					r.FormValue("deadyear1"),
					r.FormValue("periode1"),
					r.FormValue("country_id1"),
					r.FormValue("mdp"))
	}

	if r.Method == http.MethodPost {
		destroyTable(r.FormValue("name2"))
	}



	tpl, err := template.ParseFiles(filepath.Join(templateDir, "update.html"))
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	tpl.Execute(w, nil)
}