package service

import (
	"log"
	"testing"
)

func TestOtakudesuGetOngoing(t *testing.T) {
	svc := NewOtakudesuScraper()

	data, err := svc.GetOngoing()
	if err != nil {
		t.Fatal(err)
	}

	log.Printf("Total Data: %d\n", len(data))
	log.Printf("Data: %+v\n", data)
}

func TestOtakudesuGetOngoingAll(t *testing.T) {
	svc := NewOtakudesuScraper()

	data, err := svc.GetOngoingAll()
	if err != nil {
		t.Fatal(err)
	}

	log.Printf("Total Data: %d\n", len(data))
}

func TestOtakudesuGetComplete(t *testing.T) {
	svc := NewOtakudesuScraper()

	data, err := svc.GetComplete()
	if err != nil {
		t.Fatal(err)
	}

	log.Printf("Total Data: %d\n", len(data))
	log.Printf("Data: %+v\n", data)
}

func TestOtakudesuGetCompleteAll(t *testing.T) {
	svc := NewOtakudesuScraper()

	data, err := svc.GetCompleteAll()
	if err != nil {
		t.Fatal(err)
	}

	log.Printf("Total Data: %d\n", len(data))
}
