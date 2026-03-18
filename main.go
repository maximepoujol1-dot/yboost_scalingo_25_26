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

	db.Preload("Viking").Find(&pays)
	db.Preload("Country").Find(&vikings)

	
	db.Preload("Vikings").Find(&event)

	for _, e := range event {
		fmt.Println("Event:", e.Name_event)
		for _, v := range e.Vikings {
			fmt.Println(" - Viking:", v.Name)
		}
	}

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


	port := os.Getenv("PORT")
	if port == "" {
		port = "8081" 
	}

	fmt.Println("Server started at :" + port)

	err := http.ListenAndServe(":"+port, nil)
	if err != nil {
		fmt.Println("Server error:", err)
	}
}


