package main

import (
	"fmt"

	"./models"
)

func main() {
	var db models.EngineDatabase
	if err := db.Load("engines.xml"); err != nil {
		panic(err)
	}
	var rkt models.RocketModel
	if err := rkt.Load("rocket.xml"); err != nil {
		panic(err)
	}
	if err := rkt.LoadEngines(db); err != nil {
		panic(err)
	}
	fmt.Println(rkt)
}
