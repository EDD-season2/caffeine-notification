package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// SubscriptionRequest represents reuqest Dto for notification subscription request
type SubscriptionRequest struct {
	Token string `json:"token" binding:"required"`
}

type SubscriptionController struct {
	service NotificationService
}

// HandleShopSubscribe handles shop push notification subscription request
func (ctrl *SubscriptionController) HandleShopSubscribe(c *gin.Context) {
	var req SubscriptionRequest
	c.ShouldBindJSON(&req)
	shopID := c.Param("shopId")
	ctrl.service.subscribeShop(shopID, req.Token)
	c.String(http.StatusOK, "")
}

// HandleCustomerSubscribe handles shop push notification subscription request
func (ctrl *SubscriptionController) HandleCustomerSubscribe(c *gin.Context) {
	var req SubscriptionRequest
	c.ShouldBindJSON(&req)
	customerID := c.Param("customerId")
	ctrl.service.subscribeCustomer(customerID, req.Token)
	c.String(http.StatusOK, "")
}

func NewSubscriptionController(service NotificationService) *SubscriptionController {
	ctrl := new(SubscriptionController)
	ctrl.service = service
	return ctrl
}
