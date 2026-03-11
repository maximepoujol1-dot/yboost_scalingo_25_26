package main

import "gorm.io/gorm"

type Note struct {
	gorm.Model
	viking_id uint   `gorm:"not null;index"`
	Content   string `gorm:"not null"`
}

//struct API

type Viking struct{
	Viking_id  	int 	`json:"viking_id"`
	Image      	string 	`json:"image"`
	Name       	string 	`json:"name"`
	Burthyear 	int 	`json:"burthyear "`
	Deadyear 	int 	`json:"deathyear"`
	Periode 	string 	`json:"periode"`
}

type Event struct{	
	Viking_id 	int 	`json:"viking_id"`
	Name_event 	string 	`json:"viking_id"`
	Event_id 	int 	`json:"viking_id"`
}

type Country struct{
	Country_id 		int 	`json:"viking_id"`
	Country_name 	string 	`json:"viking_id"`
	Flag_country 	string 	`json:"viking_id"`
}