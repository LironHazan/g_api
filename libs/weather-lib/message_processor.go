package weather_lib

import (
	"context"
	"g_api/libs/weather-lib/ent"
	"g_api/libs/weather-lib/internal"
	"log"
)

// O2M Forcast:Weather, construct the forecast separately, n forecast = num of days
// construct each weather by iterating each day and start populating the hours, weather unit is per hour
// and the relation to the forecast is per day.
// Eventually could query all hours of a specific date, all dates, all hours in temp > 20c etc..

func ProcessMsg(msg Message, client *ent.Client, ctx context.Context) {
	// Get msg and transform data and push to table
	log.Println("processing message: ")
	forecasts, err := internal.ConstructForecasts(msg.Msg.Value)
	if err != nil {
		log.Println(err)
	}
	for _, f := range forecasts {
		// push new forecast to pg
		_, err := client.Forecast.Create().
			SetDate(f.Date).
			SetCountry(f.Country).
			SetRegion(f.Region).
			SetTemp(f.Temp).
			SetIcon(f.Icon).
			SetLocaltime(f.Localtime).
			Save(ctx)

		if err != nil {
			log.Println(err)
		}
		//todo: add hours iteration and push weather to pg

	}
}
