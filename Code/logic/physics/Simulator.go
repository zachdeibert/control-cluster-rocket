package physics

import "errors"

const (
	gravity    = 9.81
	airDensity = 1.225e-6
)

// Simulate runs the simulation on the rocket using the Improved Euler's method
func (sim *Simulation) Simulate() {
	if sim.TimeStep <= 0 {
		panic(errors.New("Invalid simulation time step"))
	}
	sim.SimulatedSteps = make([]SimulationStep, 0)
	prev := sim.Basis
	massCurve := sim.CalculateMassCurve()
	thrustCurve := sim.CalculateMassCurve()
	for t := sim.Basis.Time; prev.Velocity >= 0; t = t + sim.TimeStep {
		var next SimulationStep
		next.Time = t
		force := -gravity                                                                                                 // Gravity
		force = force - float64(sim.Rocket.CD)*airDensity*prev.Velocity*prev.Velocity/2*float64(sim.Rocket.ReferenceArea) // Drag
		force = force + LinearInterpolation2(t, thrustCurve)                                                              // Thrust
		mass := LinearInterpolation2(t, massCurve)
		next.Acceleration = force / mass
		next.Velocity = prev.Velocity + (next.Acceleration+prev.Acceleration)/2*sim.TimeStep
		next.Altitude = prev.Altitude + (next.Velocity+prev.Velocity)/2*sim.TimeStep
		sim.SimulatedSteps = append(sim.SimulatedSteps, next)
		prev = next
	}
}
