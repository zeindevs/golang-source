// package customer holds our aggrets that combines many entities into
// a full object
package customer

import (
	"errors"

	"github.com/google/uuid"
	"github.com/zeindevs/tavern"
)

var (
	ErrInvalidPerson = errors.New("a customer has to have a valid name")
)

type Customer struct {
	// person is the root entity of customer
	// which means person.ID is the main identification
	person       *tavern.Person
	products     []*tavern.Item
	transactions []tavern.Transaction
}

func NewCustomer(name string) (Customer, error) {
	return Customer{
		person: &tavern.Person{
			ID:   uuid.New(),
			Name: name,
			Age:  24,
		},
	}, nil
}

func (c *Customer) GetID() uuid.UUID {
	return c.person.ID
}

func (c *Customer) SetID(id uuid.UUID) {
	c.person.ID = id
}

func (c *Customer) GetName() string {
	return c.person.Name
}

func (c *Customer) SetName(name string) {
	c.person.Name = name
}
