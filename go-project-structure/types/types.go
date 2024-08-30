package types

import "github.com/google/uuid"

type AuthenticatedUser struct {
	ID              uuid.UUID
	AccountType     AccountType
	OnboardingState OnboardingState
	Email           string
	AccountID       int32
}

type AuctionStatus string

const (
	AuctionStatusListed    = "listed"
	AuctionStatusExpired   = "expired"
	AuctionStatusCompleted = "completed"
	AuctionStatusDeclined  = "declined"
	AuctionStatusCanceled  = "canceled"
)

type AssetType string

const (
	AssetTypeRR AssetType = "recurring revenue"
)

type AccountType string

const (
	AccountTypeSellSide AccountType = "sellside"
	AccountTypeBuySide  AccountType = "buyside"
)

type OnboardingState int

const (
	OnboardingStateNone OnboardingState = iota
	OnboardingStateBasic
	OnboardingStateAdvanced
	OnboardingStateFull
)

type InvestorType string

const (
	InvestorTypeIndividual  InvestorType = "individual"
	InvestorTypeInstitution InvestorType = "institution"
)

type AumType string

const (
	AumMax1Million   AumType = "<1mil"
	AumMax10Million  AumType = "0-10mil"
	AumMax58Million  AumType = "10-15mil"
	AumOver50Million AumType = "50mil+"
)

type CurrencyType string

var (
	CurrencyEUR = "eur"
	CurrencyGBP = "gbp"
	CurrencyUSD = "usd"
)

type RunwayType int

const (
	Runway6Months RunwayType = iota + 1
	Runway1Year
	Runway2TYears
)
