package main

import (
	"../../common"
	"./bus"
	"./models"
)

func main() {
	profile, err := models.LoadFlightProfile("test.csv")
	if err != nil {
		panic(err)
	}
	client := common.ConnectMQTT()
	if err = bus.Transmit(profile, client, []bus.TransmitConfig{
		bus.TransmitConfig{Topic: "/kinematics/position/altitude", Field: "Altitude Feet", Scale: 304.8},
		bus.TransmitConfig{Topic: "/kinematics/velocity/altitude", Field: "y-Velocity Meters / Sec", Scale: 1000},
		bus.TransmitConfig{Topic: "/kinematics/acceleration/altitude", Field: "y-Acceleration Total Meters/sec/sec", Scale: 1000},
	}, "Time"); err != nil {
		panic(err)
	}
}
