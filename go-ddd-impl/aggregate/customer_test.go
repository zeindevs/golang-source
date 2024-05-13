package aggregate_test

import (
	"errors"
	"testing"

	"ddd-impl/aggregate"
)

func TestCustomer(t *testing.T) {
	type testCase struct {
		test        string
		name        string
		expextedErr error
	}

	testCases := []testCase{
		{
			test:        "Empty name validation",
			name:        "",
			expextedErr: aggregate.ErrInvalidPerson,
		},
		{
			test:        "Valid name",
			name:        "Percy Bolmer",
			expextedErr: nil,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			_, err := aggregate.NewCustomer(tc.name)
			if !errors.Is(err, tc.expextedErr) {
				t.Errorf(`expected error %v, got %v`, tc.expextedErr, err)
			}
		})
	}
}
