package physics

import "errors"

// Interpolable represents a data point that other data can be interpolated from
type Interpolable struct {
	x float64
	y float64
}

// LinearInterpolation calculates the real position of a point by linear
// interpolating it between two data points
func LinearInterpolation(x float64, data []Interpolable) float64 {
	if x <= data[0].x {
		return data[0].y
	}
	if last := data[len(data)-1]; x >= last.x {
		return last.y
	}
	for i, p := range data {
		if p.x == x {
			return p.y
		} else if p.x > x {
			return p.y + (p.y-data[i-1].y)/(p.x-data[i-1].x)*(p.x-x)
		}
	}
	panic(errors.New("wtf"))
}

// LinearInterpolation2 calculates the real position of a point by linear
// interpolating it between two data points with multiple curves added together
func LinearInterpolation2(x float64, data [][]Interpolable) float64 {
	var sum float64
	for _, series := range data {
		sum = sum + LinearInterpolation(x, series)
	}
	return sum
}
