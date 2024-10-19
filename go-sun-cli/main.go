package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/fatih/color"
)

var (
	API_KEY = ""
)

type Weather struct {
	Location struct {
		Name            string `json:"name"`
		Country         string `json:"country"`
		LocaltimeEpoach int64  `json:"localtime_epoch"`
	} `json:"location"`
	Current struct {
		TempC     float64 `json:"temp_c"`
		Condition struct {
			Text string `json:"text"`
		} `json:"condition"`
	} `json:"current"`
	Forecast struct {
		Forecastday []struct {
			Hour []struct {
				TimeEpoach int64   `json:"time_epoch"`
				TempC      float64 `json:"temp_c"`
				Condition  struct {
					Text string `json:"text"`
				} `json:"condition"`
				ChanceOfRain float64 `json:"chance_of_rain"`
			} `json:"hour"`
		} `json:"forecastday"`
	} `json:"forecast"`
}

func main() {
	q := ""
	if len(os.Args) >= 2 {
		q = os.Args[1]
	}
	res, err := http.Get("http://api.weatherapi.com/v1/forecast.json?key=" + API_KEY + "&q=" + q + "&days=1&aqi=no&alerts=no")
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()

	if res.StatusCode != 200 {
		log.Fatal("Weather API not available")
	}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}

	var weather Weather
	if err := json.Unmarshal(body, &weather); err != nil {
		log.Fatal(err)
	}

	// if err := os.WriteFile("data.json", body, os.ModePerm); err != nil {
	//    log.Fatal(err)
	// }

	location, current, hours := weather.Location, weather.Current, weather.Forecast.Forecastday[0].Hour

	fmt.Printf("%s, %s, %s: %.0fC, %s\n",
		time.Unix(location.LocaltimeEpoach, 0).Local(),
		location.Name,
		location.Country,
		current.TempC,
		current.Condition.Text,
	)
	for _, hour := range hours {
		date := time.Unix(hour.TimeEpoach, 0)

		if date.Before(time.Now()) {
			continue
		}

		message := fmt.Sprintf("%s - %.0fC, %.0f, %s\n",
			date.Format("15:04"),
			hour.TempC,
			hour.ChanceOfRain,
			hour.Condition.Text,
		)

		if hour.ChanceOfRain < 40 {
			fmt.Print(message)
		} else {
			color.Blue(message)
		}
	}
}
