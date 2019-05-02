package bus

import (
	"fmt"
	"strconv"
	"time"

	"../models"
	mqtt "github.com/eclipse/paho.mqtt.golang"
)

// TransmitConfig describes how to send the flight profile over the bus
type TransmitConfig struct {
	Topic string
	Field string
	Scale float64
}

// Transmit sends an entire profile over the mqtt bus
func Transmit(profile models.FlightProfile, client mqtt.Client, config []TransmitConfig, timeField string) error {
	for _, cfg := range config {
		if !profile.ContainsField(cfg.Field) {
			return fmt.Errorf("Unable to find field %s", cfg.Field)
		}
	}
	if !profile.ContainsField(timeField) {
		return fmt.Errorf("Unable to find field %s", timeField)
	}
	startTime := time.Now().UnixNano()
	end := profile.NumDataPoints()
	for i := 0; i < end; i = i + 1 {
		targetTime := profile.GetData(timeField, i)
		targetNanos := startTime + int64(targetTime*1000*1000*1000)
		now := time.Now().UnixNano()
		if now < targetNanos {
			time.Sleep(time.Duration(targetNanos-now) * time.Nanosecond)
		}
		for _, cfg := range config {
			data := profile.GetData(cfg.Field, i) * cfg.Scale
			if token := client.Publish(cfg.Topic, 0, false, strconv.FormatFloat(data, 'f', -1, 64)); token.Wait() && token.Error() != nil {
				return token.Error()
			}
		}
	}
	return nil
}
