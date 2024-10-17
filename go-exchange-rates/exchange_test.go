package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func mockGetExchangeRatesUSD(date string) map[string]float64 {
	return map[string]float64{
		"USD": 1.0, "CAD": 0.74, "EUR": 1.11, "AUD": 0.67, "SRD": 0.034, "MNT": 0.00030,
	}
}

func mockGetExchangeRatesUSD_RaturnNil(date string) map[string]float64 {
	return nil
}

func TestCalcExchangeOnDate(t *testing.T) {
	type testCase struct {
		value float64
		base  string
		conv  string
		date  string

		expected float64
		err      error
	}

	t.Run("valid exchange rates", func(t *testing.T) {
		tests := []testCase{
			{value: 100.0, base: "USD", conv: "EUR", date: "2010-12-15", expected: 111},
			{value: 100.0, base: "CAD", conv: "AUD", date: "2010-12-15", expected: 90.54},
			{value: 100.0, base: "SRD", conv: "MNT", date: "2010-12-15", expected: 0.88},
		}
		for _, test := range tests {
			actual, err := CalcExchangeDate(test.value, test.base, test.conv, test.date, mockGetExchangeRatesUSD)
			assert.NoError(t, err)
			assert.Equal(t, test.expected, actual)
		}
	})

	t.Run("nil value returned", func(t *testing.T) {
		_, err := CalcExchangeDate(100, "AUD", "CAD", "2010-01-01", mockGetExchangeRatesUSD_RaturnNil)
		assert.ErrorIs(t, err, conversionError)
	})
}
