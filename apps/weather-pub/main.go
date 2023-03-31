package main

import (
	"fmt"
	"g_api/libs/weather-lib"
	"github.com/confluentinc/confluent-kafka-go/kafka"
	"os"

	"time"
)

func main() {
	_producer, err := kafka.NewProducer(&kafka.ConfigMap{
		"bootstrap.servers": "localhost:9092", //todo: env variable
	})
	if err != nil {
		fmt.Printf("failed to create producer: %s\n", err)
		return
	}
	defer _producer.Close()
	ticker := time.NewTicker(12 * time.Hour)
	weather_lib.CollectForecasts(_producer, os.Getenv("API_KEY"))
	for range ticker.C {
		weather_lib.CollectForecasts(_producer, os.Getenv("API_KEY"))
	}
}
