package handlers

import (
	"github.com/gabrielyea/simple-go-api/lib/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

// @Summary Get users
// @Description Get all users
// @Produce  json
// @Success 200 {object} services.MockUser
// @Router /users [get]
// GetUsers returns a list of users
func GetUsers(c *gin.Context) {
	user := services.GetUser()
	c.JSON(http.StatusOK, user)
}
