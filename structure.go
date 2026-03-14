package main

type Viking struct{
	Viking_id uint   `gorm:"primaryKey;column:viking_id"`
    Name      string `gorm:"column:name"`
    Image     string `gorm:"column:image"`
    Burthyear int    `gorm:"column:burthyear"`
    Deadyear  int    `gorm:"column:deadyear"`
    Periode   string `gorm:"column:periode"`
    CountryID uint    `gorm:"column:country_id"`           
    Country   Country `gorm:"foreignKey:CountryID;references:Country_id"` // relation
}

type Event struct{	
	Event_id   	uint    `gorm:"primaryKey;column:event_id"`
	Name_event 	string 	`gorm:"column:name_event"`
	VikingID  uint 		`gorm:"column:viking_id"`
	Viking    Viking `gorm:"foreignKey:VikingID"` // Relation
}

type Country struct{

	Country_id   uint     `gorm:"primaryKey;column:country_id"`
    Country_name string   `gorm:"column:country_name"`
	Flag_country string   `gorm:"column:flag_country"`
}