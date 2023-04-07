package weather_lib

import (
	"fmt"
	"g_api/libs/weather-lib/internal"
	"github.com/confluentinc/confluent-kafka-go/kafka"
	"sync"
)

type Message struct {
	Msg   *kafka.Message
	Topic string
}

type MessageHandler func(msg Message)

func ConsumeForecast(topic string, consumer *kafka.Consumer, onSuccess MessageHandler, wg *sync.WaitGroup) error {
	defer wg.Done()

	fmt.Printf("Subscribed to topic: [  %s  ] \n", topic)

	for {
		msg, err := consumer.ReadMessage(-1)
		if err != nil {
			fmt.Printf("error while consuming message: %v\n", err)
			return err
		}
		fmt.Printf("operating on received message in topic %s\n", topic)
		onSuccess(Message{Msg: msg, Topic: topic})
	}
}

func SubscribeToForecastUpdates(consumer *kafka.Consumer, onSuccess MessageHandler, onError ...func(error)) {
	err := consumer.SubscribeTopics([]string{"tel_aviv", "los_angeles", "paris"}, nil)
	if err != nil {
		onError[0](err)
	}
	var wg sync.WaitGroup
	for _, topic := range internal.RegionToTopic() {
		wg.Add(1)
		go ConsumeForecast(topic, consumer, onSuccess, &wg)
	}
	wg.Wait()
}
