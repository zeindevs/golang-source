package data

import (
	"context"

	"github.com/zeindevs/go-project-structure/db"
	"github.com/zeindevs/go-project-structure/types"
)

type Account struct {
	ID              int64 `bun:"id,pk,autoincrement"`
	AccountType     types.AccountType
	OnboardingState types.OnboardingState
}

func CreateAccount(accountType types.AccountType) (*Account, error) {
	account := &Account{
		AccountType:     accountType,
		OnboardingState: types.OnboardingStateBasic,
	}
	_, err := db.Bun.NewInsert().Model(account).Exec(context.Background())
	return account, err
}
