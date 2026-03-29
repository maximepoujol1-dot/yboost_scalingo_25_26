package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func addTable(name string, image string, burth string, dead string, periode string, country string, mdp string){
	
    if mdp != os.Getenv("Mdp") {
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
				CountryID: id_country,
			}
					
	if err := db.Where(Viking{Name: name}).FirstOrCreate(&new).Error; err != nil {
		if strings.Contains(err.Error(), "viking_pkey") {
			if seqErr := realignVikingSequence(); seqErr != nil {
				fmt.Printf("⚠️ Erreur insertion et realignement sequence impossible: %v | %v\n", err, seqErr)
				return
			}

			if retryErr := db.Where(Viking{Name: name}).FirstOrCreate(&new).Error; retryErr != nil {
				fmt.Printf("⚠️ Erreur insertion apres realignement sequence: %v\n", retryErr)
				return
			}

			fmt.Println("✓ Insertion reussie apres realignement de sequence")
			return
		}

		fmt.Printf("⚠️ Erreur insertion viking: %v\n", err)
	}
}

func updateTable(id string, name string, image string, burth string, dead string, periode string, country string, mdp string){
	
    if mdp != os.Getenv("Mdp") {
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
	vikingID, _ := strconv.Atoi(id)
	
	updateData := map[string]interface{}{
		"name":       name,
		"image":      image,
		"burthyear":  burthInt,
		"deadyear":   deathInt,
		"periode":    periode,
		"country_id": id_country,
	}

	if err := db.Model(&Viking{}).Where("viking_id = ?", vikingID).Updates(updateData).Error; err != nil {
		fmt.Printf("⚠️ Erreur update viking: %v\n", err)
	}
}

func destroyTable(name string, mdp string){

	if mdp != os.Getenv("Mdp") {
        fmt.Println("⚠️ Erreur : Mot de passe incorrect ou variable d'env non définie")
        return

    } 
	if err := db.Where("name = ?", name).Delete(&Viking{}).Error; err != nil {
		fmt.Printf("⚠️ Erreur suppression viking: %v\n", err)
		return
	}

	if err := realignVikingSequence(); err != nil {
		fmt.Printf("⚠️ Erreur realignement sequence apres suppression: %v\n", err)
	}
}
