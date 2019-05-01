package flight

const (
	initialAccuracy      = 20
	pllBufferSize        = 100
	calcSafetyFactor     = 1.3 // safety factor on time spent doing calculations
	maxFlightTime        = 30  // seconds to apogee
	timeResolution       = 0.1 // seconds
	gradientDescentDelta = 0.000001
	maxOptimalTimeChange = 0.0001            // maximum change in optimal fire time over time to allow firing
	minOptimizationTime  = 200 * 1000 * 1000 // maximum time (in nanoseconds) the optimal calculations have to be the same for before firing
)
