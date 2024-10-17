package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

	_ "github.com/joho/godotenv/autoload"
)

var API_KEY = os.Getenv("API_KEY")

type Response struct {
	Rates map[string]float64
}

const urlTemplate = "https://openexchangerates.org/api/historical/%s.json?app_id=%s"

func getExchangeRatesUSD(date string) map[string]float64 {
	url := fmt.Sprintf(urlTemplate, date, API_KEY)
	resp, err := http.Get(url)
	if err != nil {
		log.Printf("Failed to get response: %v", err)
		return nil
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Printf("Failed to read response body: %v", err)
		return nil
	}

	if err := os.WriteFile(fmt.Sprintf("data/%s.json", date), body, os.ModePerm); err != nil {
		log.Printf("Failed to write to json: %v", err)
	}

	var data Response
	if err := json.Unmarshal(body, &data); err != nil {
		log.Printf("Failed to unmarshal JSON: %v", err)
		return nil
	}

	return data.Rates
}

func getLocalExchangeRatesUSD(date string) map[string]float64 {
	file, err := os.ReadFile(fmt.Sprintf("data/%s.json", date))
	if err != nil {
		log.Printf("Failed to load rates, file not found")
		return nil
	}
	var data Response
	if err := json.Unmarshal(file, &data); err != nil {
		log.Printf("Failed to unmarshal JSON: %v", err)
		return nil
	}

	return data.Rates
}
