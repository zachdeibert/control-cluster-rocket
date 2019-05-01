package flight

import (
	"time"

	"../models"
)

type optimizedComponent struct {
	Component models.RocketComponent
	Fire      func(models.RocketComponent)
	Fired     bool
}

func (comp *optimizedComponent) UpdateTime(burnTime float64) {
	if !comp.Fired && burnTime < float64(time.Now().UnixNano()) {
		comp.Fired = true
		comp.Fire(comp.Component)
	}
}
