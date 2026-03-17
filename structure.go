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

func (Viking) TableName() string {
    return "Viking" // le nom exact de ta table Supabase
}

type Event struct{	
	Event_id   	uint    `gorm:"primaryKey;column:event_id"`
	Name_event 	string 	`gorm:"column:name_event"`
}

func (Event) TableName() string {
    return "Event" // le nom exact de ta table Supabase
}

type Country struct{

	Country_id   uint     `gorm:"primaryKey;column:country_id"`
    Country_name string   `gorm:"column:country_name"`
	Flag_country string   `gorm:"column:flag_country"`
    Population string       `gorm:"column:population"`
    Politique string `gorm:"column:politique"`
}

func (Country) TableName() string {
    return "Country" // le nom exact de ta table Supabase
}