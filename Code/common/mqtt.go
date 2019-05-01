package common

import (
	"math/rand"
	"time"

	MQTT "github.com/eclipse/paho.mqtt.golang"
)

const alphabet = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

// ConnectMQTT connects to the MQTT server running locally on the rocket
func ConnectMQTT() MQTT.Client {
	rand.Seed(time.Now().UTC().UnixNano())
	opts := MQTT.NewClientOptions().AddBroker("tcp://127.0.0.1:1883")
	cs := make([]byte, 23-len("rocket"))
	for i := range cs {
		cs[i] = alphabet[rand.Int63()%int64(len(alphabet))]
	}
	opts.SetClientID("rocket" + string(cs))
	c := MQTT.NewClient(opts)
	if token := c.Connect(); token.Wait() && token.Error() != nil {
		panic(token.Error())
	}
	return c
}
