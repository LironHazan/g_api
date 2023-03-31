package main

import (
	"fmt"
	weather_lib "g_api/libs/weather-lib"
	"github.com/confluentinc/confluent-kafka-go/kafka"
)

func main() {
	_consumer, err := kafka.NewConsumer(&kafka.ConfigMap{
		"bootstrap.servers": "localhost:9092", //todo: env variable
		"group.id":          "my-group-id",
		"auto.offset.reset": "earliest",
	})
	if err != nil {
		fmt.Printf("failed to create consumer: %s\n", err)
		return
	}
	defer _consumer.Close()
	weather_lib.SubscribeToForecastUpdates(_consumer)
}
