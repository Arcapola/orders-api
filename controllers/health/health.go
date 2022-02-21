// Package health exposes health check handlers.
package health

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Controller handles health check requests.
type Controller struct{}

// Ping godoc
// @Summary  ping function
// @Description  run health-check ping request
// @Tags         health
// @Accept       json
// @Produce      json
// @Success      200  {string}  pong
// @Router       /health/ping [get]
// Respond to user ping request.
func (p *Controller) Ping(c *gin.Context) {
	c.String(http.StatusOK, "pong")
}
