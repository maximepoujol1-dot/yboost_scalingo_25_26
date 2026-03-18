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
    Vikings     []Viking  `gorm:"many2many:EventViking;foreignKey:Event_id;joinForeignKey:Event_id;references:Viking_id;joinReferences:Viking_id"`
}

func (Event) TableName() string {
    return "Event" // le nom exact de ta table Supabase
}

type EventViking struct {

    Event_id uint `gorm:"column:event_id"`
    Event Event `gorm:"foreignKey:Event_id;references:Event_id"` // relation

    Viking_id uint   `gorm:"column:viking_id"`
    Viking *Viking `gorm:"foreignKey:Viking_id;references:Viking_id"` // relation
    
}

func (EventViking) TableName() string {
    return "EventViking" 
}



type Country struct{

	Country_id   uint     `gorm:"primaryKey;column:country_id"`
    Country_name string   `gorm:"column:country_name"`
	Flag_country string   `gorm:"column:flag_country"`
    Population string       `gorm:"column:population"`
    Politique string `gorm:"column:politique"`

    VikingID uint   `gorm:"column:viking_id"` 
    Viking *Viking `gorm:"foreignKey:VikingID;references:Viking_id"` // relation

}

func (Country) TableName() string {
    return "Country" // le nom exact de ta table Supabase
}