package main

import (
	"io"
	"log/slog"
	"net/http"
	"os"
	"strings"

	_ "github.com/joho/godotenv/autoload"
)

func main() {
	url := "https://api.api2convert.com/v2/jobs"

	payload := strings.NewReader(`{
    "input": [{
      "type": "remote",
      "source": "https://example-files.online-convert.com/raster%20image/jpg/example_small.jpg"
    }],
    "conversion": [{
      "category": "category",
      "target": "png",
      "options": {}
    }]
  }`)

	client := &http.Client{}
	req, err := http.NewRequest("POST", url, payload)
	if err != nil {
		slog.Warn("post request failed", "error", err.Error())
		return
	}

	req.Header.Set("x-oc-api-key", os.Getenv("API_KEY"))
	req.Header.Set("Content-Type", "application/json")

	res, err := client.Do(req)
	if err != nil {
		slog.Warn("do request failed", "error", err.Error())
		return
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		slog.Warn("read body failed", "error", err.Error())
	}

	slog.Info("response", "data", body)
}
