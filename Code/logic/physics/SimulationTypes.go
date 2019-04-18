package physics

import (
	"fmt"

	"../models"
)

// SimulationStep represents a single instant of all of the values
type SimulationStep struct {
	Time         float64
	Acceleration float64
	Velocity     float64
	Altitude     float64
}

// MotorBurn describes an event in which one of the motors is burned
type MotorBurn struct {
	Time  float64
	Model models.RocketComponent
}

// Simulation describes a simulation that should be run
type Simulation struct {
	Basis          SimulationStep
	Burns          []MotorBurn
	Rocket         models.RocketModel
	TimeStep       float64
	SimulatedSteps []SimulationStep
}

// String converts the SimulationStep to a string
func (step SimulationStep) String() string {
	return fmt.Sprintf("(t = %f, a = %f, v = %f, x = %f)", step.Time, step.Acceleration, step.Velocity, step.Altitude)
}

// String converts the MotorBurn to a string
func (burn MotorBurn) String() string {
	return fmt.Sprintf("Burning motor at t=%f: {%s}", burn.Time, burn.Model)
}

// String converts the Simulation to a string
func (sim Simulation) String() string {
	return fmt.Sprintf("Simulation starting at %s with dt=%f and burns %s: %s (rocket = %s)", sim.Basis, sim.TimeStep, sim.Burns, sim.SimulatedSteps, sim.Rocket)
}

// Apogee finds the apogee altitude from the simulation
func (sim Simulation) Apogee() float64 {
	var last float64
	for i := len(sim.SimulatedSteps) - 1; i >= 0; i = i - 1 {
		if sim.SimulatedSteps[i].Altitude < last {
			return last
		}
		last = sim.SimulatedSteps[i].Altitude
	}
	return last
}
