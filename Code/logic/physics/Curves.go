package physics

// CalculateMassCurve creates a mass curve for the rocket
func (sim Simulation) CalculateMassCurve() [][]Interpolable {
	curve := make([][]Interpolable, len(sim.Rocket.Components))
	for i, comp := range sim.Rocket.Components {
		curve[i] = make([]Interpolable, len(comp.Model.Data))
		for j, data := range comp.Model.Data {
			curve[i][j] = Interpolable{
				x: float64(data.Time),
				y: float64(data.Mass),
			}
		}
	}
	return curve
}

// CalculateThrustCurve creates a thrust curve for the rocket
func (sim Simulation) CalculateThrustCurve() [][]Interpolable {
	curve := make([][]Interpolable, len(sim.Rocket.Components))
	for i, comp := range sim.Rocket.Components {
		curve[i] = make([]Interpolable, len(comp.Model.Data))
		for j, data := range comp.Model.Data {
			curve[i][j] = Interpolable{
				x: float64(data.Time),
				y: float64(data.Force),
			}
		}
	}
	return curve
}
