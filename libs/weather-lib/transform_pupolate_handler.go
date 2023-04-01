package weather_lib

import (
	"context"
	"g_api/libs/weather-lib/ent"
	"log"
)

func ProcessMsg(msg Message, client *ent.Client, ctx context.Context) {
	// Get msg and transform data and push to table per topic
	if err := client.Schema.Create(ctx); err != nil {
		log.Fatal(err)
	}
}
