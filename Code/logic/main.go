package main

import (
	"fmt"

	"./models"
)

func main() {
	var db models.EngineDatabase
	db.Load("engines.xml")
	fmt.Println(db)
}
