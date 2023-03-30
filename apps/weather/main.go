package main

import (
	"g_api/libs/weather-lib"
	"time"
)

func main() {
	ticker := time.NewTicker(12 * time.Hour)
	weather_lib.TriggerForcastPipline()
	for range ticker.C {
		weather_lib.TriggerForcastPipline()
	}
}
