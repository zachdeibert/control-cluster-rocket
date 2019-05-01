package flight

import "../physics"

type optimizationStep struct {
	Optimizer Optimizer
	Position  physics.SimulationStep
	Accuracy  float64
}

func (step *optimizationStep) calculate(burnTimes []float64) float64 {
	sim := physics.Simulation{
		Basis:    step.Position,
		Burns:    make([]physics.MotorBurn, len(burnTimes)),
		Rocket:   step.Optimizer.Rocket,
		TimeStep: step.Accuracy,
	}
	for i, time := range burnTimes {
		sim.Burns[i] = physics.MotorBurn{
			Time:  time,
			Model: step.Optimizer.Rocket.Components[i],
		}
	}
	sim.Simulate()
	return sim.Apogee()
}
