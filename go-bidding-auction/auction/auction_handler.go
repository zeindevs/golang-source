package main

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"sync"
	"time"
)

type BiddingServiceResponse struct {
	AdID     string `json:"ad_id"`
	BidPrice int    `json:"bid_price"`
}

func auctionHandler(w http.ResponseWriter, r *http.Request) {
	// Extract AdPlacementId from request
	adPlacementId := r.URL.Query().Get("ad_placement_id")
	log.Println("Request received for AdPlacementId:", adPlacementId)

	biddingServices := []string{
		"http://bidding:80801/bid",
		"http://bidding-2:8001/bid",
		"http://bidding-3:8001/bid",
	}

	// create channels to receive bid responses and errors
	bidResponses := make(chan *BiddingServiceResponse, len(biddingServices))
	bidErrors := make(chan error, len(biddingServices))

	var wg sync.WaitGroup

	for _, serviceURL := range biddingServices {
		wg.Add(1)
		ctx, _ := context.WithTimeout(r.Context(), time.Millisecond*200)
		go func(url string) {
			defer wg.Done()
			bidResponse, err := callBiddingService(ctx, url)
			if err != nil {
				bidErrors <- err
				return
			}
			log.Println("Received bid response from: ", url, "Ad Id", bidResponse.AdID, "Bid Price", bidResponse.BidPrice)
			bidResponses <- bidResponse
		}(serviceURL)
	}

	log.Println("Waiting for all responses")
	wg.Wait()
	log.Println("All responses received")

	close(bidResponses)
	close(bidErrors)

	// iterate over the bidResponses channel and select the highest bid
	var bestBid *BiddingServiceResponse

	for bid := range bidResponses {
		if bestBid == nil || bid.BidPrice > bestBid.BidPrice {
			bestBid = bid
		}
	}

	if bestBid == nil {
		w.WriteHeader(http.StatusNoContent)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(bestBid)
}

func callBiddingService(ctx context.Context, url string) (*BiddingServiceResponse, error) {
	client := &http.Client{
		Timeout: time.Millisecond * 200,
	}

	resp, err := client.Get(url)
	if err != nil {
		return nil, err
	}

	var bidResponse BiddingServiceResponse
	if err := json.NewDecoder(resp.Body).Decode(&bidResponse); err != nil {
		return nil, err
	}

	return &bidResponse, nil
}
