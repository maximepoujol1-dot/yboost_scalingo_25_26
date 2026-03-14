package main

import (

	"html/template"
	"net/http"
	"os"
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
	

	tpl, err := template.ParseFiles(filepath.Join(templateDir, "index.html"))
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	tpl.Execute(w, vikings)
}


func homeHandler2(w http.ResponseWriter, r *http.Request) {

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
	

	tpl, err := template.ParseFiles(filepath.Join(templateDir, "pays.html"))
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	tpl.Execute(w, pays)
}