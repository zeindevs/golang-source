package data

import (
	"context"
	"time"

	"github.com/zeindevs/go-project-structure/db"
)

type Analyses struct {
	ID                int64     `bun:"id,pk,autoincrement"`
	SellSideCompanyID int64     `json:"sellSideCompanyID"`
	AvgMRR            float64   `json:"avgMRR"`
	AvgNetChurn       float64   `json:"avgNetChurn"`
	AvgNetGrowth      float64   `json:"avgNetGrowth"`
	Runway            int       `json:"runway"`
	CreatedAt         time.Time `json:"createdAt"`
}

type TradingLimit struct {
	ID         int64     `bun:"id,pk,autoincrement"`
	AnalysesID int64     `json:"analysesID"`
	Amount     float64   `json:"amount"`
	CreatedAt  time.Time `json:"createdAt"`
}

type CreateAnalysisParams struct {
	AvgMRR       float64 `json:"avgMRR"`
	AvgNetChurn  float64 `json:"avgNetChurn"`
	AvgNetGrowth float64 `json:"avgNetGrowth"`
	Runway       int     `json:"runway"`
	TradingLimit float64 `json:"tradingLimit"`
}

func CreateAnlysis(companyID int64, params CreateAnalysisParams) error {
	ctx := context.Background()
	tx, err := db.Bun.Begin()
	if err != nil {
		return err
	}
	a := Analyses{
		SellSideCompanyID: companyID,
		AvgMRR:            params.AvgMRR,
		AvgNetChurn:       params.AvgNetChurn,
		AvgNetGrowth:      params.AvgNetGrowth,
		Runway:            params.Runway,
	}
	_, err = db.Bun.NewInsert().Model(&a).Exec(ctx)
	if err != nil {
		return err
	}
	tr := TradingLimit{
		AnalysesID: a.ID,
		Amount:     params.TradingLimit,
	}
	_, err = db.Bun.NewInsert().Model(&tr).Exec(ctx)
	if err != nil {
		return err
	}
	return tx.Commit()
}

func GetTradingLimit(accountID int64) (*TradingLimit, error) {
	var tradingLimit TradingLimit
	err := db.Bun.NewSelect().
		Model(&tradingLimit).
		Join("join analyses as ana on ana.id = trading_limit.analyses_id").
		Join("join sell_side_companies as ssc on ssc.id = ana.sell_side_company_id").
		Join("join companies as com on com.id = ssc.company_id").
		Join("join accounts as acc on acc.id = com.account_id").
		Where("acc.id = ?", accountID).
		Scan(context.Background())
	return &tradingLimit, err
}
