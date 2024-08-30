package data

import "github.com/zeindevs/go-project-structure/types"

type Company struct {
	ID int64
}

func GetFundingRequests(acccountID int64) (any, error) {
	return nil, nil
}

func CreateFundingRequest(user *types.AuthenticatedUser, cfr CreateFundingRequestParams) (*Funding, error) {
	return &Funding{}, nil
}

func GetSellSideCompanyByAccountID(accountID int64) (*Company, error) {
	return nil, nil
}
