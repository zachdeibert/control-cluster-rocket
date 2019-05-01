package main

import (
	"fmt"
	"strconv"
	"time"

	"../common"
	"./flight"
	"./models"
	mqtt "github.com/eclipse/paho.mqtt.golang"
)

var (
	client            mqtt.Client
	launchTime        int64 = -1
	kinematics              = make([]float64, 3)
	kinematicsUpdated       = make([]bool, 3)
	logic             *flight.Optimizer
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
	client = common.ConnectMQTT()
	logic = flight.NewOptimizer(rkt, fireMotor)
	client.Subscribe("/control/launch", 0, func(_ mqtt.Client, msg mqtt.Message) {
		launchTime = time.Now().UnixNano()
	})
	client.Subscribe("/kinematics/position/altitude", 0, func(_ mqtt.Client, msg mqtt.Message) {
		kinematicHandler(0, msg)
	})
	client.Subscribe("/kinematics/velocity/altitude", 0, func(_ mqtt.Client, msg mqtt.Message) {
		kinematicHandler(1, msg)
	})
	client.Subscribe("/kinematics/acceleration/altitude", 0, func(_ mqtt.Client, msg mqtt.Message) {
		kinematicHandler(2, msg)
	})
}

func kinematicHandler(idx int, msg mqtt.Message) {
	if parsed, err := strconv.ParseFloat(string(msg.Payload()), 64); err != nil {
		fmt.Printf("Unable to read kinematic broadcast: %s\n", err)
	} else {
		kinematics[idx] = parsed
		if launchTime > 0 {
			kinematicsUpdated[idx] = true
			good := true
			for _, val := range kinematicsUpdated {
				if !val {
					good = false
					break
				}
			}
			if good {
				logic.UpdateData(float64(time.Now().UnixNano()-launchTime)/1000000000, kinematics[2], kinematics[1], kinematics[0])
				for i := range kinematicsUpdated {
					kinematicsUpdated[i] = false
				}
			}
		}
	}
}

func fireMotor(motor models.RocketComponent) {
	if motor.Ignition.Type == models.BoosterEngine {
		client.Publish(fmt.Sprintf("/boosters/%d/fire", motor.Ignition.ID), 2, false, common.BoosterFireMagic)
	}
}
