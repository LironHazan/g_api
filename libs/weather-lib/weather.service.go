package weather_lib

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

const apiEndpoint = "https://api.weatherapi.com/v1/forecast.json"

type Result = interface{}

func queryForcast(region string) (Result, error) {
	req, err := http.NewRequest("GET", apiEndpoint, nil)
	if err != nil {
		return nil, err
	}

	// set the query parameters
	q := req.URL.Query()
	q.Add("key", os.Getenv("API_KEY"))
	q.Add("q", region)
	q.Add("days", "10")
	q.Add("aqi", "no")
	q.Add("alerts", "no")
	req.URL.RawQuery = q.Encode()

	// make the request and parse the response
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var result interface{}
	if err := json.Unmarshal(body, &result); err != nil {
		return nil, err
	}
	// Consider typing the json and use an assertion
	//_result, ok := result.(Result)
	//if !ok {
	//	// ...
	//}
	return result, nil
}

func pushForcast(forcast any) {
	fmt.Println(forcast)
}

func TriggerForcastPipline() {

	// Phase1
	// call queryForcast
	// push to topics [[IL] , [UK], [US]]
	// consume + transform
	// push to pg

	// API + UI
	// model gql schema
	// generate API
	// UI consumer (like the iphone app)

	regions := [3]string{"Tel-Aviv", "Paris", "Los Angeles"}

	for _, r := range regions {
		forcast, err := queryForcast(r)
		if err != nil {
			pushForcast(forcast)
		}
	}

}
