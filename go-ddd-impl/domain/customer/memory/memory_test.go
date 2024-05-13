package memory

import (
	"errors"
	"testing"

	"ddd-impl/aggregate"
	"ddd-impl/domain/customer"

	"github.com/google/uuid"
)

func TestMemory(t *testing.T) {
	type testCase struct {
		name        string
		id          uuid.UUID
		expectedErr error
	}

	cust, err := aggregate.NewCustomer("percy")
	if err != nil {
		t.Fatal(err)
	}

	id := cust.GetID()

	repo := MemoryRepository{
		customers: map[uuid.UUID]aggregate.Customer{
			id: cust,
		},
	}

	testCases := []testCase{
		{
			name:        "no custome by id",
			id:          uuid.MustParse("ffd2de91-3de5-47ca-bcf2-024f0847869e"),
			expectedErr: customer.ErrCustomerNotFound,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			_, err := repo.Get(tc.id)

			if !errors.Is(err, tc.expectedErr) {
				t.Errorf("expected error %v, got %v", tc.expectedErr, err)
			}
		})
	}
}
