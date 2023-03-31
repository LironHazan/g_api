package weather_lib

import (
	"fmt"
	"github.com/confluentinc/confluent-kafka-go/kafka"
)

func ConsumeForecast(topic string, consumer *kafka.Consumer, messages chan []byte) {
	err := consumer.SubscribeTopics([]string{topic}, nil)
	if err != nil {
		panic(err)
	}

	for {
		msg, err := consumer.ReadMessage(-1)
		if err != nil {
			fmt.Printf("error while consuming message: %v\n", err)
			continue
		}
		fmt.Printf("received message: %v\n", string(msg.Value))
		messages <- msg.Value
	}
}

func SubscribeToForecastUpdates(consumer *kafka.Consumer) {
	ch := make(chan []byte)

	for _, topic := range RegionToTopic() {
		go func(t string) {
			ConsumeForecast(t, consumer, ch)
		}(topic)
	}
	for msg := range ch {
		fmt.Println(string(msg))
	}
}
