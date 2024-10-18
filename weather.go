package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

const (
	apiKey = "93307f7c93dfe9b37bd56cc196ede40c"
	city   = "Jakarta"
	units  = "metric"
)

// Struct to parse the JSON response
type WeatherResponse struct {
	List []struct {
		Dt   int64 `json:"dt"`
		Main struct {
			Temp float64 `json:"temp"`
		} `json:"main"`
		DtTxt string `json:"dt_txt"`
	} `json:"list"`
}

func main() {
	// Create the API URL for 5-day weather forecast
	url := fmt.Sprintf("https://api.openweathermap.org/data/2.5/forecast?q=%s&appid=%s&units=%s", city, apiKey, units)

	// Make the HTTP GET request
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer resp.Body.Close()

	// Parse the JSON response
	var weatherResponse WeatherResponse
	if err := json.NewDecoder(resp.Body).Decode(&weatherResponse); err != nil {
		fmt.Println("Error decoding JSON:", err)
		return
	}

	forecastedDays := map[string]bool{} // To avoid duplicate days

	fmt.Println("Weather Forecast:")
	for _, weather := range weatherResponse.List {
		// Convert the UNIX timestamp to a readable date format
		forecastTime := time.Unix(weather.Dt, 0)
		dayFormat := forecastTime.Format("Mon, 02 Jan 2006")

		// Ensure we only show one forecast per day
		if _, exists := forecastedDays[dayFormat]; !exists {
			fmt.Printf("%s: %.2f°C\n", dayFormat, weather.Main.Temp)
			forecastedDays[dayFormat] = true
		}
	}

}
