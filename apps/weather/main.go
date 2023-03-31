package main

import (
	"g_api/libs/weather-lib"

	"time"
)

func main() {
	//brokers := []string{"localhost:9092"}
	//producer, err := producer.NewProducer(brokers)

	ticker := time.NewTicker(12 * time.Hour)
	weather_lib.TriggerForcastPipline()
	for range ticker.C {
		weather_lib.TriggerForcastPipline()
	}
}
