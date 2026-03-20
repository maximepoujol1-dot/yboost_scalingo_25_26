package main

type Viking struct {
	Viking_id uint    `gorm:"primaryKey;column:viking_id"`
	Name      string  `gorm:"column:name"`
	Image     string  `gorm:"column:image"`
	Burthyear int     `gorm:"column:burthyear"`
	Deadyear  int     `gorm:"column:deadyear"`
	Periode   string  `gorm:"column:periode"`
	CountryID uint    `gorm:"column:country_id"`
	Country   Country `gorm:"foreignKey:CountryID;references:Country_id"`
	Events    []Event `gorm:"many2many:EventViking;foreignKey:Viking_id;joinForeignKey:viking_id;References:Event_id;joinReferences:event_id"`
}

func (Viking) TableName() string { return "Viking" }

type Event struct {
	Event_id   uint     `gorm:"primaryKey;column:event_id"`
	Name_event string   `gorm:"column:name_event"`
	Vikings    []Viking `gorm:"many2many:EventViking;foreignKey:Event_id;joinForeignKey:event_id;References:Viking_id;joinReferences:viking_id"`
}

func (Event) TableName() string { return "Event" }

type EventViking struct {
	Event_id  uint `gorm:"column:event_id"`
	Viking_id uint `gorm:"column:viking_id"`
}

func (EventViking) TableName() string { return "EventViking" }

type Country struct {
	Country_id   uint     `gorm:"primaryKey;column:country_id"`
	Country_name string   `gorm:"column:country_name"`
	Flag_country string   `gorm:"column:flag_country"`
	Population   string   `gorm:"column:population"`
	Politique    string   `gorm:"column:politique"`
	Vikings      []Viking `gorm:"foreignKey:CountryID;references:Country_id"`
}

func (Country) TableName() string { return "Country" }