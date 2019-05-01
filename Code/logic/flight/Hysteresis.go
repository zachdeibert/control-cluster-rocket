package flight

import (
	"math"
	"time"
)

type hysteresisFunc struct {
	BaseFunction  func(float64)
	MaxDxDt       float64 // max change in value over time (in nanoseconds)
	MinGoodTime   int64   // minimum time (in nanoseconds) the value must follow the maximum dx/dt curve for
	goodTimeStart int64   // nanoseconds
	lastTime      int64   // nanoseconds
	last          float64
}

func (h *hysteresisFunc) callback(x float64) {
	t := time.Now().UnixNano()
	dxDt := (x - h.last) / float64(t-h.lastTime)
	if math.Abs(dxDt) > h.MaxDxDt || h.lastTime < 0 {
		h.goodTimeStart = t
		h.lastTime = t
	}
	h.last = x
	if t-h.goodTimeStart > h.MinGoodTime {
		h.BaseFunction(x)
	}
}

func applyHysteresis(base func(float64), maxDxDt float64, minGoodTime int64) func(float64) {
	h := hysteresisFunc{
		BaseFunction:  base,
		MaxDxDt:       maxDxDt,
		MinGoodTime:   minGoodTime,
		goodTimeStart: 0,
		lastTime:      -1,
		last:          0,
	}
	return h.callback
}
