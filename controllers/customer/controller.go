// Package customer exposes CRUD API over Customer model
package customer

// Controller handles operations over orders.
type Controller struct{}

// NewController initializes a controller for order model.
func NewController() *Controller {
	return &Controller{}
}
