package data

import (
	"context"

	"github.com/uptrace/bun"
)

type Metric struct{}

func GetMetric(ctx context.Context, db *bun.DB, v any) error {
	// TODO: GetMetric
	return nil
}
