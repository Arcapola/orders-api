package models

import (
	"github.com/gofrs/uuid"
)

// Customer is a model representing customer.
type Customer struct {
	ID      uuid.UUID `json:"id" example:"f3537d5c-cf15-48cc-b309-7a51312a574f" format:"uuid"`
	Name    string    `json:"name" example:"A Company Ltd." format:"string"`
	Address string    `json:"address" example:"Street 1234, 01234 City, Country" format:"string"`
}

// AddCustomer is a limited creation version of customer object.
type AddCustomer struct {
	Name    string `json:"name" example:"A Company Ltd." format:"string"`
	Address string `json:"address" example:"Street 1234, 01234 City, Country" format:"string"`
}

// Validate whether customer creation request is correct.
func (a AddCustomer) Validate() error {
	return nil
}

// Insert newly created customer object into database.
func (c *Customer) Insert() (uuid.UUID, error) {
	genID, err := uuid.NewV4()
	if err != nil {
		return uuid.Nil, err
	}

	c.ID = genID
	return genID, nil
}

// FindCustomerByID finds customer by ID and returns it.
func FindCustomerByID(id uuid.UUID) (Customer, error) {
	// TODO: Implement DB lookup.
	return Customer{
		ID:      id,
		Name:    "Test Customer",
		Address: "Street 12/A, 01234 City, Country",
	}, nil
}
