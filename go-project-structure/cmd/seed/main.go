package main

import (
	"fmt"
	"log"

	"github.com/google/uuid"
	"github.com/zeindevs/go-project-structure/data"
	"github.com/zeindevs/go-project-structure/db"
	"github.com/zeindevs/go-project-structure/types"
)

func main() {
	db.Init()
	seed := true

	params := data.OnboardingParams{AccountType: types.AccountTypeSellSide,
		PhoneNumber:     "00000000",
		CompanyPosition: "CTO",
		CompanyName:     "Foobar inc",
		Website:         "https://website.com",
		Country:         "Belgium",
		SellSideParams: data.SellSideOnboardingParams{
			RevenueType: "recurring",
			CompanySize: "20-80",
		},
	}

	userID := uuid.MustParse("")
	if err := data.OnboardUser(userID, params); err != nil {
		log.Fatal(err)
	}

	details, err := data.GetUserDetails(userID)
	if err != nil {
		log.Fatal(err)
	}

	sellSideCompany, err := data.GetSellSideCompanyByAccountID(details.AccountID)
	if err != nil {
		log.Fatal(err)
	}

	if err := SeedAnalysis(sellSideCompany.ID); err != nil {
		log.Fatal(err)
	}

	if seed {
		authUser := &types.AuthenticatedUser{
			ID:        userID,
			AccountID: int32(details.AccountID),
		}
		cfr := data.CreateFundingRequestParams{
			Amount:      500000,
			Currency:    types.CurrencyEUR,
			Reason:      "Need money now",
			Period:      12,
			MarketPlace: 0,
		}
		_, err := data.CreteFundingRequest(authUser, cfr)
		if err != nil {
			log.Fatal(err)
		}
		cfr = data.CreateFundingRequestParams{
			Amount:      250000,
			Currency:    types.CurrencyEUR,
			Reason:      "Need money now",
			Period:      12,
			MarketPlace: 0,
		}
		fundingResp, err := data.CreateFundingRequest(authUser, cfr)
		if err != nil {
			log.Fatal(err)
		}
		if err := data.SetAuctionStatus(fundingResp.AuctionID, authUser, types.AuctionStatusCompleted); err != nil {
			log.Fatal(err)
		}
		cfr = data.CreateFundingRequestParams{
			Amount:      500000,
			Currency:    types.CurrencyEUR,
			Reason:      "Need money now",
			Period:      12,
			MarketPlace: 0,
		}
		fundingResp, err = data.CreateFundingRequest(authUser, cfr)
		if err != nil {
			log.Fatal(err)
		}
		if err := data.SetAuctionStatus(fundingResp.AuctionID, authUser, types.AuctionStatusDeclined); err != nil {
			log.Fatal(err)
		}
	}

	fundings, err := data.GetFundingRequests(details.AccountID)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(fundings)
}

func SeedAnalysis(companyID int64) error {
	params := data.CreateAnalysisParams{
		AvgMRR:       2300.15,
		AvgNetChurn:  600.43,
		AvgNetGrowth: 300.15,
		Runway:       4,
		TradingLimit: 1000000000,
	}
	return data.CreateAnalysis(companyID, params)
}
