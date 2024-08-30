package data

import "github.com/zeindevs/go-project-structure/types"

type Funding struct {
	AuctionID int64
}

func CreteFundingRequest(authUser *types.AuthenticatedUser, cfr CreateFundingRequestParams) (*Funding, error) {
	return nil, nil
}

func CreateAnalysis(companyID int64, params CreateAnalysisParams) error {
	return nil
}
