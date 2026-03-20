package main

var vikings []Viking
var pays []Country
var event []Event

func createTable(){
	
}

func destroyTable(){
	
}

func updateTable(){
	
}

func loadTable(){
	db.Preload("Viking").Find(&pays)
	db.Preload("Country").Find(&vikings)
	db.Preload("Vikings").Find(&event)
}