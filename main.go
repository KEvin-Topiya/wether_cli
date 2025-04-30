package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"os"

	"github.com/fatih/color"
)

// Structs for geocoding
type GeocodingResponse struct {
	Results []struct {
		Latitude  float64 `json:"latitude"`
		Longitude float64 `json:"longitude"`
		Name      string  `json:"name"`
		Country   string  `json:"country"`
	} `json:"results"`
}

// Structs for weather
type WeatherResponse struct {
	CurrentWeather struct {
		Temperature float64 `json:"temperature"`
		WindSpeed   float64 `json:"windspeed"`
		WeatherCode int     `json:"weathercode"`
		Time        string  `json:"time"`
	} `json:"current_weather"`
}

// Weather codes (simple interpretation)
var weatherConditions = map[int]string{
	0:  "Clear sky",
	1:  "Mainly clear",
	2:  "Partly cloudy",
	3:  "Overcast",
	45: "Fog",
	48: "Depositing rime fog",
	51: "Light drizzle",
	53: "Moderate drizzle",
	55: "Dense drizzle",
	61: "Slight rain",
	63: "Moderate rain",
	65: "Heavy rain",
	71: "Slight snow",
	73: "Moderate snow",
	75: "Heavy snow",
	80: "Rain showers",
	95: "Thunderstorm",
}

func getCoordinates(city string) (float64, float64, string, error) {
	encoded := url.QueryEscape(city)
	url := fmt.Sprintf("https://geocoding-api.open-meteo.com/v1/search?name=%s&count=1", encoded)

	resp, err := http.Get(url)
	if err != nil {
		return 0, 0, "", err
	}
	defer resp.Body.Close()

	var geo GeocodingResponse
	if err := json.NewDecoder(resp.Body).Decode(&geo); err != nil {
		return 0, 0, "", err
	}

	if len(geo.Results) == 0 {
		return 0, 0, "", fmt.Errorf("city not found")
	}

	return geo.Results[0].Latitude, geo.Results[0].Longitude, fmt.Sprintf("%s, %s", geo.Results[0].Name, geo.Results[0].Country), nil
}

func getWeather(lat, lon float64) (WeatherResponse, error) {
	url := fmt.Sprintf("https://api.open-meteo.com/v1/forecast?latitude=%.4f&longitude=%.4f&current_weather=true", lat, lon)
	resp, err := http.Get(url)
	if err != nil {
		return WeatherResponse{}, err
	}
	defer resp.Body.Close()

	var weather WeatherResponse
	err = json.NewDecoder(resp.Body).Decode(&weather)
	return weather, err
}

func main() {
	if len(os.Args) < 2 {
		color.Yellow("Usage: go run weather.go [city]")
		return
	}
	city := os.Args[1]

	color.Cyan("🔍 Looking up coordinates for %s...", city)
	lat, lon, place, err := getCoordinates(city)
	if err != nil {
		color.Red("Error: %v", err)
		return
	}

	color.Green("📍 Location: %s (%.2f, %.2f)", place, lat, lon)

	weather, err := getWeather(lat, lon)
	if err != nil {
		color.Red("Error fetching weather: %v", err)
		return
	}

	color.Blue("\n🌤️ Weather in %s", place)
	fmt.Printf("🌡️  Temperature: %.1f°C\n", weather.CurrentWeather.Temperature)
	fmt.Printf("💨 Wind Speed: %.1f km/h\n", weather.CurrentWeather.WindSpeed)
	fmt.Printf("🌈 Conditions: %s\n", weatherConditions[weather.CurrentWeather.WeatherCode])
	fmt.Printf("⏰ Time: %s\n", weather.CurrentWeather.Time)

}
