package controller

import (
	"net/http"
	"pwaV3/config"
	"pwaV3/models"
	"time"

	"github.com/gin-gonic/gin"
)

type OrderController struct{}

func (oc *OrderController) CreateOrder(c *gin.Context) {
	var request struct {
		Status               string             `json:"status"`
		RestaurantLocationID uint               `json:"restaurant_location_id"`
		CustomerID           uint               `json:"customer_id"`
		Items                []models.OrderItem `json:"items"`
		Payment              uint               `json:"payment_id"`
	}

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error bind": err.Error()})
		return
	}

	// Create the order
	order := models.Order{
		Status:               request.Status,
		RestaurantLocationID: request.RestaurantLocationID,
		OrderItems:           request.Items,
		CustomerID:           request.CustomerID,
		Date:                 time.Now(),
		PaymentID:            request.Payment,
	}

	if err := config.DB.Create(&order).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error create": err.Error()})
		return
	}

	// Iterate over the order items and assign the order_id
	for i := range order.OrderItems {
		order.OrderItems[i].OrderID = order.ID
	}

	// Save the order items
	if err := config.DB.Create(&order.OrderItems).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error save": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"order": order})
}
