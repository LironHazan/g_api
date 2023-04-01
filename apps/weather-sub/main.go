package main

import (
	"context"
	"fmt"
	weather_lib "g_api/libs/weather-lib"
	"g_api/libs/weather-lib/db"
	"github.com/confluentinc/confluent-kafka-go/kafka"
	"log"
)

func main() {
	client := db.Open("postgresql://user:password@127.0.0.1/database") //todo env vars
	ctx := context.Background()
	err := handleForecast(func(msg weather_lib.Message) {
		weather_lib.ProcessMsg(msg, client, ctx)
	})
	if err != nil {
		panic(err)
	}
}

func handleForecast(onSuccess weather_lib.MessageHandler) error {
	_consumer, err := initConsumer()
	if err != nil {
		fmt.Printf("failed to create consumer: %s\n", err)
		return err
	}
	defer _consumer.Close()
	weather_lib.SubscribeToForecastUpdates(_consumer, onSuccess, onError)
	return nil
}

func onError(err error) {
	log.Printf("failed: %s\n", err)
}

func initConsumer() (*kafka.Consumer, error) {
	return kafka.NewConsumer(&kafka.ConfigMap{
		"bootstrap.servers": "localhost:9092", //todo: env variable
		"group.id":          "my-group-id",
		"auto.offset.reset": "earliest",
	})
}
