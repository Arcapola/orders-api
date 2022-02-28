package models

import (
	"time"

	"github.com/gofrs/uuid"
)

// Order is a model representing order.
type Order struct {
	ID        uuid.UUID  `json:"id" example:"f3537d5c-cf15-48cc-b309-7a51312a574f" format:"uuid"`
	Customer  Customer   `json:"customer"`
	Items     []MenuItem `json:"items"`
	CreatedAt time.Time  `json:"created_at"`
}

// AddOrder is a limited creation version of order object.
type AddOrder struct {
	CustomerID uuid.UUID  `json:"customer_id"`
	Items      []MenuItem `json:"items"`
}

// Validate whether order creation request is correct.
func (a AddOrder) Validate() error {
	return nil
}

// Insert newly created order object into database.
func (o *Order) Insert() (uuid.UUID, error) {
	genID, err := uuid.NewV4()
	if err != nil {
		return uuid.Nil, err
	}

	o.ID = genID
	return genID, nil
}

// MenuItem is a model representing an item on the menu.
type MenuItem struct {
	Name      string  `json:"name" example:"Food" format:"string"`
	Count     uint    `json:"count" example:"1" format:"uint"`
	UnitPrice float64 `json:"unit_price" example:"3.99" format:"float64"`
}
