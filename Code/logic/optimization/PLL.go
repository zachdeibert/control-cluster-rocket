package optimization

import (
	"errors"
	"time"
)

type pllQueue struct {
	Time    int64
	Payload interface{}
}

type pllRecord struct {
	Time     int64
	Accuracy float64
}

// PLLOptimizeLoop optimizes the accuracy of a calculation by maximizing the
// amount of processor time used between input events with the constraint of
// being done with calculations before the next input event arrives
func PLLOptimizeLoop(input chan interface{}, accuracy float64, bufferSize int, safetyFactor float64, loop func(interface{}, float64)) {
	if bufferSize < 1 {
		panic(errors.New("Invalid buffer size"))
	}
	queue := make(chan pllQueue)
	go func() {
		times := make([]int64, bufferSize)
		ptr := -1
		filled := 0
		lastTime := int64(0)
		for {
			select {
			case data := <-input:
				var entry pllQueue
				if ptr < 0 {
					entry = pllQueue{
						Time:    -1,
						Payload: data,
					}
				} else {
					ptr = (ptr + 1) % bufferSize
					times[ptr] = time.Now().UnixNano() - lastTime
					if filled < bufferSize {
						filled = filled + 1
					}
					var sum float64
					for _, time := range times {
						sum = sum + float64(time)
					}
					entry = pllQueue{
						Time:    int64(sum / float64(filled)),
						Payload: data,
					}
				}
				lastTime = time.Now().UnixNano()
				go func() {
					queue <- entry
				}()
				break
			}
		}
	}()
	accuracyMap := make([]pllRecord, bufferSize)
	ptr := 0
	filled := 0
	for {
		e := <-queue
		thisAccuracy := accuracy
		if e.Time > 0 {
			var timeSum int64
			var accuracySum float64
			for i := 0; i < filled; i = i + 1 {
				timeSum = timeSum + accuracyMap[i].Time
				accuracySum = accuracySum + accuracyMap[i].Accuracy
			}
			if timeSum > 0 {
				thisAccuracy = accuracySum * float64(e.Time) / float64(timeSum)
			}
		}
		thisAccuracy = thisAccuracy / safetyFactor
		startTime := time.Now().UnixNano()
		loop(e.Payload, thisAccuracy)
		endTime := time.Now().UnixNano()
		accuracyMap[ptr] = pllRecord{
			Time:     endTime - startTime,
			Accuracy: thisAccuracy,
		}
		ptr = (ptr + 1) % bufferSize
		if filled < bufferSize {
			filled = filled + 1
		}
	}
}
