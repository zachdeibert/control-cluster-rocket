package flight

import (
	"math"

	"../models"
	"../optimization"
	"../physics"
)

// Optimizer is a class that can optimize the times that the motors are lit
type Optimizer struct {
	Rocket      models.RocketModel
	DataChannel chan interface{}
	Motors      []func(float64)
}

// NewOptimizer creates a new Optimizer object
func NewOptimizer(rocket models.RocketModel, fire func(models.RocketComponent)) *Optimizer {
	opt := Optimizer{
		Rocket:      rocket,
		DataChannel: make(chan interface{}),
		Motors:      make([]func(float64), len(rocket.Components)),
	}
	for i, comp := range rocket.Components {
		c := optimizedComponent{
			Component: comp,
			Fire:      fire,
			Fired:     false,
		}
		opt.Motors[i] = applyHysteresis(c.UpdateTime, maxOptimalTimeChange, minOptimizationTime)
	}
	go optimization.PLLOptimizeLoop(opt.DataChannel, initialAccuracy, pllBufferSize, calcSafetyFactor, opt.recalculate)
	return &opt
}

// UpdateData tells the optimizer the new location of the rocket from sensors
func (opt *Optimizer) UpdateData(time float64, acceleration float64, velocity float64, altitude float64) {
	opt.DataChannel <- physics.SimulationStep{
		Time:         time,
		Acceleration: acceleration,
		Velocity:     velocity,
		Altitude:     altitude,
	}
}

func (opt *Optimizer) recalculate(_pos interface{}, accuracy float64) {
	pos := _pos.(physics.SimulationStep)
	basis := make([]float64, len(opt.Rocket.Components))
	for i := range basis {
		basis[i] = maxFlightTime / 2
	}
	step := optimizationStep{
		Optimizer: *opt,
		Position:  pos,
		Accuracy:  accuracy,
	}
	bestTimes := optimization.GradientDescent(basis, step.calculate, gradientDescentDelta, maxFlightTime/4, int(math.Ceil(math.Log2(maxFlightTime/timeResolution))))
	for i, val := range bestTimes {
		opt.Motors[i](val)
	}
}
