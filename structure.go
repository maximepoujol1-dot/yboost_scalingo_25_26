package main





type Viking struct{
	viking_id 	uint   `gorm:"not null;index"`
	Image     	string `gorm:"image"`
	Name      	string `gorm:"name"`
	Burthyear 	int    `gorm:"burthyear"`
	Deadyear  	int    `gorm:"deathyear"`
	Periode   	string `gorm:"periode"`
}

type Event struct{	
	Viking_id  	int    `gorm:"viking_id"`
	Name_event 	string `gorm:"name_event"`
	Event_id   	int    `gorm:"event_id"`
}

type Country struct{
	Country_id   int    `gorm:"country_id"`
	Country_name string `gorm:"country_name"`
	Flag_country string `gorm:"flag_country"`
}