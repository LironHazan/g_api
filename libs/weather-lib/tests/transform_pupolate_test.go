package tests

import (
	"g_api/libs/weather-lib/internal"
	_ "github.com/mattn/go-sqlite3"
	"io"
	"os"
	"testing"
	"time"
)

func loadData(filename string) ([]byte, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	content, err := io.ReadAll(file)
	if err != nil {
		return nil, err
	}

	return content, nil
}

func TestParseLocationData(t *testing.T) {
	data := map[string]interface{}{
		"location": map[string]interface{}{
			"name":            "Tel Aviv-Yafo",
			"region":          "Tel Aviv",
			"country":         "Israel",
			"lat":             32.07,
			"lon":             34.76,
			"tz_id":           "Asia/Jerusalem",
			"localtime_epoch": 1680762757,
			"localtime":       "2023-04-06 9:32",
		},
	}

	location, err := internal.ParseLocationData(data)
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}

	expectedRegion := "Tel Aviv"
	if location.Region != expectedRegion {
		t.Errorf("Expected region %q, but got %q", expectedRegion, location.Region)
	}

	expectedCountry := "Israel"
	if location.Country != expectedCountry {
		t.Errorf("Expected country %q, but got %q", expectedCountry, location.Country)
	}

	expectedLocalTime := time.Date(2023, 4, 6, 9, 32, 0, 0, time.UTC)
	if location.Localtime != expectedLocalTime {
		t.Errorf("Expected localtime %q, but got %q", expectedLocalTime, location.Localtime)
	}
}

//func Test_constructForecast(t *testing.T) {
//	data, err := loadData("data.json")
//	forecasts, err := internal.ConstructForecasts(data)
//	if err != nil {
//		return
//	}
//	fmt.Println(forecasts)
//}
