package main

type Viking struct{
	Viking_id 	uint   `gorm:"primaryKey"`
	Image     	string 
	Name      	string 
	Burthyear 	int    
	Deadyear  	int    
	Periode   	string 
	CountryID  uint      
	Country    Country `gorm:"foreignKey:CountryID"` // Relation avec Country
	Events     []Event `gorm:"foreignKey:VikingID"` // Relation  avec Event
}

type Event struct{	
	Event_id   	uint    `gorm:"primaryKey"`
	Name_event 	string 
	VikingID  uint 
	Viking    Viking `gorm:"foreignKey:VikingID"` // Relation Many-to-One
	
	
}

type Country struct{
	Country_id   int    `gorm:"primaryKey"`
	Country_name string
	Flag_country string 
	Vikings     []Viking `gorm:"foreignKey:CountryID"` // Relation One-to-Many
}