// Package order exposes CRUD API over Order model
package order

// Controller handles operations over orders.
type Controller struct{}

// NewController initializes a controller for order model.
func NewController() *Controller {
	return &Controller{}
}
