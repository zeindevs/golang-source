package customer_test

import (
	"errors"
	"testing"

	"github.com/zeindevs/tavern/domain/customer"
)

func TestCreateNewCustomer(t *testing.T) {
	type testCase struct {
		test        string
		name        string
		expectedErr error
	}

	testCases := []testCase{
		{
			test:        "Empty name validation",
			name:        "",
			expectedErr: customer.ErrInvalidPerson,
		},
		{
			test:        "valid name",
			name:        "Percy Bolmer",
			expectedErr: nil,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.test, func(t *testing.T) {
			_, err := customer.NewCustomer(tc.name)

			if !errors.Is(err, tc.expectedErr) {
				t.Errorf("expected error %v, got %v", tc.expectedErr, err)
			}
		})
	}
}
