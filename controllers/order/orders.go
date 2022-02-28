package order

import (
	"net/http"
	"time"

	"github.com/arcapola/orders-api/httputil"
	"github.com/arcapola/orders-api/models"
	"github.com/gin-gonic/gin"
)

// AddOrder godoc
// @Summary	Add an order
// @Description	Add an order by JSON and store it in-memory (TODO)
// @Tags	orders
// @Accept	json
// @Produce	json
// @Param	order	body		models.AddOrder	true	"Add order"
// @Success	200	{object}	models.Order
// @Failure	400	{object}	httputil.HTTPError
// @Failure	500	{object}	httputil.HTTPError
// @Router	/order	[post]
// Create a new order for customer.
func (c *Controller) AddOrder(ctx *gin.Context) {
	var addOrder models.AddOrder
	if err := ctx.ShouldBindJSON(&addOrder); err != nil {
		httputil.NewError(ctx, http.StatusBadRequest, err)
		return
	}

	if err := addOrder.Validate(); err != nil {
		httputil.NewError(ctx, http.StatusBadRequest, err)
		return
	}

	customer, err := models.FindCustomerByID(addOrder.CustomerID)
	if err != nil {
		httputil.NewError(ctx, http.StatusBadRequest, err)
		return
	}

	order := models.Order{
		Customer:  customer,
		Items:     addOrder.Items,
		CreatedAt: time.Now(),
	}

	lastID, err := order.Insert()
	if err != nil {
		httputil.NewError(ctx, http.StatusBadRequest, err)
		return
	}

	order.ID = lastID
	ctx.JSON(http.StatusOK, order)
}
