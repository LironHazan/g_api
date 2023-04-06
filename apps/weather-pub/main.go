package main

import (
	"fmt"
	"g_api/libs/weather-lib"
	"github.com/confluentinc/confluent-kafka-go/kafka"
	"log"
	"os"

	"time"
)

func main() {
	err := collectPeriodicForecasts()
	if err != nil {
		panic(err)
	}
}

func collectPeriodicForecasts() error {
	ticker := time.NewTicker(12 * time.Hour)
	_producer, err := initProducer()
	if err != nil {
		fmt.Printf("failed to create producer: %s\n", err)
		return err
	}
	defer _producer.Close()
	publishForecasts(_producer) // imperative error handling which I'm not sure if I like...

	for range ticker.C {
		publishForecasts(_producer)
	}
	return nil
}

func initProducer() (*kafka.Producer, error) {
	return kafka.NewProducer(&kafka.ConfigMap{
		"bootstrap.servers": "localhost:9092", //todo: env variable
	})
}

func publishForecasts(producer *kafka.Producer) {
	if err := weather_lib.PublishForecasts(producer, os.Getenv("API_KEY")); err != nil {
		log.Printf("failed to publish forecast: %s\n", err) // in real world would probably log it somewhere
	}
}
