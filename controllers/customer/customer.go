package customer

import (
	"net/http"

	"github.com/arcapola/orders-api/httputil"
	"github.com/arcapola/orders-api/models"
	"github.com/gin-gonic/gin"
)

// AddCustomer godoc
// @Summary	Add a customer
// @Description	Add a customer by JSON and store it in-memory (TODO)
// @Tags	customers
// @Accept	json
// @Produce	json
// @Param	order	body		models.AddCustomer	true	"Add customer"
// @Success	200	{object}	models.Customer
// @Failure	400	{object}	httputil.HTTPError
// @Failure	500	{object}	httputil.HTTPError
// @Router	/customer	[post]
// Create a new customer.
func (c *Controller) AddCustomer(ctx *gin.Context) {
	var addCustomer models.AddCustomer
	if err := ctx.ShouldBindJSON(&addCustomer); err != nil {
		httputil.NewError(ctx, http.StatusBadRequest, err)
		return
	}

	if err := addCustomer.Validate(); err != nil {
		httputil.NewError(ctx, http.StatusBadRequest, err)
		return
	}

	customer := models.Customer{
		Name:    addCustomer.Name,
		Address: addCustomer.Address,
	}

	lastID, err := customer.Insert()
	if err != nil {
		httputil.NewError(ctx, http.StatusBadRequest, err)
		return
	}

	customer.ID = lastID
	ctx.JSON(http.StatusOK, customer)
}
