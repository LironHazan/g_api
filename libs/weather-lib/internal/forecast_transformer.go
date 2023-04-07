package internal

import (
	"encoding/json"
	"fmt"
	"log"
	"time"
)

type Forecast struct {
	Date      time.Time `json:"date,omitempty"`
	Country   string    `json:"country,omitempty"`
	Region    string    `json:"region,omitempty"`
	Localtime time.Time `json:"localtime,omitempty"`
	Icon      string    `json:"icon,omitempty"`
	Temp      float64   `json:"temp,omitempty"`
}

type Location struct {
	Region    string
	Country   string
	Localtime time.Time
}

type Current struct {
	Icon string  `json:"icon,omitempty"`
	Temp float64 `json:"temp,omitempty"`
}

func newForecast(location Location, current Current, day string) Forecast {
	forecast := Forecast{}
	forecast.Country = location.Country
	forecast.Region = location.Region
	forecast.Localtime = location.Localtime
	forecast.Icon = current.Icon
	forecast.Temp = current.Temp
	forecast.Date, _ = time.Parse("2006-01-02 15:04", day)
	return forecast
}

func ParseLocationData(data map[string]interface{}) (Location, error) {
	locationData := data["location"].(map[string]interface{})
	region := locationData["region"].(string)
	country := locationData["country"].(string)
	localtimeStr := locationData["localtime"].(string)
	localtime, err := time.Parse("2006-01-02 15:04", localtimeStr)
	if err != nil {
		return Location{}, err
	}
	return Location{region, country, localtime}, nil
}

func ParseCurrentData(data map[string]interface{}) Current {
	currentData := data["current"].(map[string]interface{})
	icon := currentData["condition"].(map[string]interface{})["icon"].(string)
	Temp := currentData["temp_c"].(float64)

	return Current{
		Icon: icon,
		Temp: Temp,
	}
}

func ConstructForecasts(msg []byte) (f []Forecast, e error) {
	var data map[string]interface{}
	err := json.Unmarshal(msg, &data)
	if err != nil {
		log.Fatal(err)
	}

	location, err := ParseLocationData(data)
	if err != nil {
		return nil, err
	}
	current := ParseCurrentData(data)

	var forecasts []Forecast
	days, ok := data["forecast"].(map[string]interface{})["forecastday"].([]interface{})

	if !ok {
		// handle type assertion error
		return nil, fmt.Errorf("No forecast or forecastday on given data source")
	}
	for _, day := range days {
		dayMap, ok := day.(map[string]interface{})
		date := dayMap["date"].(string)
		if !ok {
			return nil, fmt.Errorf("date is not a string")
		}
		forecast := newForecast(location, current, date)
		forecasts = append(forecasts, forecast)
	}

	return forecasts, nil
}
