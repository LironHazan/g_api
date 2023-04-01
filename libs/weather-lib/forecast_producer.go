package weather_lib

import (
	"fmt"
	"github.com/confluentinc/confluent-kafka-go/kafka"
	"io"
	"log"
	"net/http"
)

func queryForecast(region string, apiKey string) ([]byte, error) {
	req, err := http.NewRequest("GET", "https://api.weatherapi.com/v1/forecast.json", nil)
	if err != nil {
		return nil, err
	}

	// set the query parameters
	q := req.URL.Query()
	q.Add("key", apiKey)
	q.Add("q", region)
	q.Add("days", "10")
	q.Add("aqi", "no")
	q.Add("alerts", "no")
	req.URL.RawQuery = q.Encode()

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("received non-OK status code: %d", res.StatusCode)
	}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	log.Printf("Received response with body size: %d", len(body))

	return body, nil
}

func pushForecast(producer *kafka.Producer, region string, forecast []byte) {
	topic := RegionToTopic()[region]
	fmt.Printf("publishing to: %v\n", topic)

	message := &kafka.Message{
		TopicPartition: kafka.TopicPartition{Topic: &topic, Partition: kafka.PartitionAny},
		Value:          forecast,
	}

	log.Printf("Received at time: %s", message.Timestamp.Format("2006-01-02 15:04:05"))

	deliveryChan := make(chan kafka.Event)
	err := producer.Produce(message, deliveryChan)
	if err != nil {
		return
	}

	e := <-deliveryChan
	m := e.(*kafka.Message)

	if m.TopicPartition.Error != nil {
		fmt.Printf("failed to deliver message: %v\n", m.TopicPartition.Error)
	} else {
		fmt.Printf("message delivered to topic %s at partition %d offset %d\n",
			*m.TopicPartition.Topic, m.TopicPartition.Partition, m.TopicPartition.Offset)
	}
}

func PublishForecasts(producer *kafka.Producer, apiKey string) error {
	for region, _ := range RegionToTopic() {
		forecast, err := queryForecast(region, apiKey)

		if err != nil {
			return err
		}

		log.Printf("Pushing body: %d", len(forecast))
		pushForecast(producer, region, forecast)
	}
	return nil
}
