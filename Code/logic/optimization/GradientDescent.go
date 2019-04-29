package optimization

import "errors"

// CalculateGradient approximates the gradient vector of a function at a given
// position
func CalculateGradient(pos []float64, objective func([]float64) float64, delta float64) []float64 {
	gradient := make([]float64, len(pos))
	for i := range pos {
		pos[i] = pos[i] - delta
		lower := objective(pos)
		pos[i] = pos[i] + 2*delta
		upper := objective(pos)
		pos[i] = pos[i] - delta
		gradient[i] = (upper - lower) / delta
	}
	return gradient
}

// GradientDescent goes down the gradient function to find a local minima of the
// objective function
func GradientDescent(basis []float64, objective func([]float64) float64, gradientDelta float64, jumpDelta float64, iterations int) []float64 {
	if iterations <= 0 {
		panic(errors.New("Invalid iteration count"))
	}
	for iter := 0; iter < iterations; iter = iter + 1 {
		gradient := CalculateGradient(basis, objective, gradientDelta)
		for i := range basis {
			basis[i] = basis[i] - gradient[i]*jumpDelta
		}
		jumpDelta = jumpDelta / 2
	}
	return basis
}
