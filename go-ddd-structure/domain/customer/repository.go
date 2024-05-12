package customer

import (
	"errors"

	"github.com/google/uuid"
)

var (
	ErrCustomerNotFound    = errors.New("the customer wa found in the repository")
	ErrFailedToAddCustomer = errors.New("failed to add the customer")
	ErrUpdateCustomer      = errors.New("failed to update the customer")
)

type Repository interface {
	Get(uuid.UUID) (Customer, error)
	Add(Customer) error
	Update(Customer) error
}
