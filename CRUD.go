package main

import (
	"fmt"
	"os"
	"strconv"
)

var vikings []Viking
var pays []Country
var event []Event

func createTable(name string, image string, burth string, dead string, periode string, country string, mdp string){
	fmt.Printf("Données reçues : %s, %s, %s\n", name, country, mdp)

    if mdp != os.Getenv("mdp") {
        fmt.Println("⚠️ Erreur : Mot de passe incorrect ou variable d'env non définie")
        return

    } 
	
	var id_country uint = 0
	if country == "Danemark" || country == "1"{
		id_country =1
	} else if country == "Norvège" || country == "2" {
		id_country =2	
	} else if country == "Suède" || country == "3"{
		id_country =3	
	} else if country == "Islande" || country == "4"{
		id_country =4	
	} 

	burthInt, _ := strconv.Atoi(burth)
    deathInt, _ := strconv.Atoi(dead)
	
	new := Viking{Name: name,
				Image: image, 
				Burthyear: burthInt, 
				Deadyear: deathInt, 
				Periode: periode, 
				CountryID: id_country}
					
	db.Create(&new)
	
}

//func destroyTable(name string){
//	db.Where("name = ?", name).Delete(&viking)
//}

//func updateTable(name string){
//	db.Model(&Viking{}).Where("name = ?", name).Update()
//}

func loadTable(){
	db.Preload("Viking").Find(&pays)
	db.Preload("Country").Find(&vikings)
	db.Preload("Vikings").Find(&event)
}