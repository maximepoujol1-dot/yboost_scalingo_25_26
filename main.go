package main

import (
	"fmt"
	"net/http"
	"os"
	"path/filepath"
	
)

var templateDir string

func main() {

	initDB()

	loadTable()

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
	http.HandleFunc("/pays", homeHandler4)
	http.HandleFunc("/event", homeHandler5)
	http.HandleFunc("/create", homeHandler6)

	port := os.Getenv("PORT")

	fmt.Println("Server started at :" + port)

	err := http.ListenAndServe(":"+port, nil)
	if err != nil {
		fmt.Println("Server error:", err)
	}
}


