package weather_lib

import (
	"context"
	"encoding/json"
	"g_api/libs/weather-lib/ent"
	"log"
)

func ProcessMsg(msg Message, client *ent.Client, ctx context.Context) {
	// Get msg and transform data and push to table per topic
	// O2M Forcast:Weather
	log.Println("processing message: ")

	var data map[string]interface{}
	err := json.Unmarshal(msg.Msg.Value, &data)
	if err != nil {
		log.Fatal(err)
	}
}
