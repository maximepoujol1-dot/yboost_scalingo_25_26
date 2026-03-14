package main

import (

	"html/template"
	"net/http"
	"strconv"
	"path/filepath"

)

var vikings []Viking
var pays []Country
var event []Event

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

	if pays == nil {
		db.Find(&pays)
	}

	tpl, err := template.ParseFiles(filepath.Join(templateDir, "pays.html"))
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	tpl.Execute(w, pays)
}

func homeHandler5(w http.ResponseWriter, r *http.Request) {
	
	if event == nil {
		db.Preload("Viking").Find(&event)
	}

	tpl, err := template.ParseFiles(filepath.Join(templateDir, "index.html"))
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	tpl.Execute(w, event)
}