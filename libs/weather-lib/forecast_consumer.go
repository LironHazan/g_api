package weather_lib

import (
	"fmt"
	"github.com/confluentinc/confluent-kafka-go/kafka"
)

type Message struct {
	Msg   []byte
	Topic string
}

type MessageHandler func(msg Message)

func ConsumeForecast(topic string, consumer *kafka.Consumer, messages chan Message) error {
	err := consumer.SubscribeTopics([]string{topic}, nil)
	if err != nil {
		return err
	}

	for {
		msg, err := consumer.ReadMessage(-1)
		if err != nil {
			fmt.Printf("error while consuming message: %v\n", err)
			return err
		}
		fmt.Printf("received message: %v\n", string(msg.Value))
		messages <- Message{Msg: msg.Value, Topic: topic}
	}
}

func SubscribeToForecastUpdates(consumer *kafka.Consumer, onSuccess MessageHandler, onError ...func(error)) {
	ch := make(chan Message) // should it be bounded?

	for _, topic := range RegionToTopic() {
		go func(t string) {
			if err := ConsumeForecast(t, consumer, ch); err != nil {
				if len(onError) > 0 {
					onError[0](err)
				}
			}
		}(topic)
	}
	for msg := range ch {
		fmt.Println(string(msg.Msg))
		onSuccess(msg)
	}
}
