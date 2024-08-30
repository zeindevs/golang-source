package data

import (
	"github.com/google/uuid"
	"github.com/zeindevs/go-project-structure/types"
)

type OnboardingParams struct {
	AccountType     types.AccountType
	PhoneNumber     string
	CompanyPosition string
	CompanyName     string
	Website         string
	Country         string
	SellSideParams  SellSideOnboardingParams
}

type SellSideOnboardingParams struct {
	RevenueType string
	CompanySize string
}

type CreateFundingRequestParams struct {
	Amount      int64
	Currency    string
	Reason      string
	Period      int
	MarketPlace int
}

func OnboardUser(userID uuid.UUID, params OnboardingParams) error {
	return nil
}

type UserDetails struct {
	AccountID int64
	Account   struct {
		OnboardingState types.OnboardingState
		AccountType     types.AccountType
	}
}

func GetUserDetails(userID uuid.UUID) (*UserDetails, error) {
	return &UserDetails{}, nil
}
