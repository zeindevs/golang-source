package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"net/http"
	"os"
)

var urlTemplate string = "https://youtube.googleapis.com/youtube/v3/channels?part=statistics&forHandle=%s&key=%s"

type ChannelStatistic struct {
	ViewCount             string
	SubscriberCount       string
	HiddenSubscriberCount bool
	VideoCount            string
}

type PageInfo struct {
	TotalResults   int
	ResultsPerPage int
}

type ChannelStatisticsItem struct {
	Kind       string
	Etag       string
	Id         string
	Statistics ChannelStatistic
}

type ChannelStatisticsApiResponse struct {
	Kind     string
	Etag     string
	PageInfo PageInfo
	Items    []ChannelStatisticsItem
}

func main() {
	var channelHandle string
	var apiKey string

	flag.StringVar(&channelHandle, "channelHandle", "", "The youtube channel handle to query.")
	flag.StringVar(&apiKey, "apiKey", "", "The Youtube Data API V3 Key to use.")
	flag.Parse()

	if apiKey == "" && channelHandle == "" {
		fmt.Println("Invalid apiKey or channelHandle provided. Aborting.")
		os.Exit(1)
	}

	fmt.Println(channelHandle, apiKey)

	url := fmt.Sprintf(urlTemplate, channelHandle, apiKey)
	resp, err := http.Get(url)
	if err != nil {
		fmt.Printf("Failed to get Youtube channel statistics: %v\n", err)
		os.Exit(1)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	// fmt.Println(len(body), string(body))

	stats := &ChannelStatisticsApiResponse{}

	err = json.Unmarshal(body, &stats)
	if err != nil {
		fmt.Println("Failed to unmarshal JSON body from response.", err)
		os.Exit(1)
	}

	if len(stats.Items) < 1 {
		fmt.Println("The channel was not found or contained no statistics.")
		os.Exit(1)
	}

	channelStats := stats.Items[0].Statistics
	fmt.Printf("--- Channel Statistics for %s ---\n", channelHandle)
	fmt.Printf("ViewCount: %s \n SubscriberCount: %s \n ViewCount: %s", channelStats.ViewCount, channelStats.SubscriberCount, channelStats.VideoCount)
}
