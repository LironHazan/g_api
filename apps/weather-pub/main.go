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
	weather_lib.PublishForecasts(_producer, os.Getenv("323c2774127c47fc9aa121442232903"))
	for range ticker.C {
		weather_lib.PublishForecasts(_producer, os.Getenv("323c2774127c47fc9aa121442232903"))
	}
}
