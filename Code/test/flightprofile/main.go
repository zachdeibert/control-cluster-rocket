package main

import (
	"fmt"

	"./models"
)

func main() {
	profile, err := models.LoadFlightProfile("test.csv")
	if err != nil {
		panic(err)
	}
	fmt.Println(profile)
}
