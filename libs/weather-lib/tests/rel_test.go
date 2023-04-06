package tests

import (
	"context"
	"fmt"
	"g_api/libs/weather-lib/ent"
	"log"
	"testing"
	"time"

	"entgo.io/ent/dialect"
	_ "github.com/mattn/go-sqlite3"
)

func Test_Weather(t *testing.T) {
	// Create an ent.Client with in-memory SQLite database.
	client, err := ent.Open(dialect.SQLite, "file:ent?mode=memory&cache=shared&_fk=1")
	if err != nil {
		log.Fatalf("failed opening connection to sqlite: %v", err)
	}
	defer client.Close()
	ctx := context.Background()
	// Run the automatic migration tool to create all schema resources.
	if err := client.Schema.Create(ctx); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}

	day, err := populateDay(client, ctx)
	if err != nil {
		log.Fatalf("failed: %v", err)
	}

	hour, err := populateHour(client, ctx)
	if err != nil {
		log.Fatalf("failed: %v", err)
	}

	day.Update().AddWeather(hour)
	fmt.Print(day.QueryWeather().All(ctx))
	fmt.Printf(" temp: %v", hour.Temp)

}

func populateDay(client *ent.Client, ctx context.Context) (*ent.Forecast, error) {
	newDate := time.Date(2023, time.March, 30, 0, 0, 0, 0, time.UTC)
	return client.Forecast.Create().
		SetDate(newDate).
		SetCountry("Israel").
		SetRegion("TelAviv").
		SetAvgTemp(10.1).
		SetMaxTemp(11.1).
		SetMinTemp(8.1).
		Save(ctx)
}

func populateHour(client *ent.Client, ctx context.Context) (*ent.Weather, error) {
	newDate := time.Date(2023, time.March, 30, 0, 0, 0, 0, time.UTC)
	return client.Weather.Create().
		SetTemp(10.1).
		SetFeelsLike(9.1).
		SetDate(newDate).
		SetTime(newDate.Hour()).
		Save(ctx)
}
