package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

// NotificationSendRequest represents request for send notification both shop and customer
type NotificationSendRequest struct {
	Message string `json:"message" binding:"required"`
}

type NotificationSendController struct {
	service NotificationService
}

func (ctrl *NotificationSendController) HandleSendShopNotification(c *gin.Context) {
	var req NotificationSendRequest
	c.ShouldBindJSON(&req)
	shopID := c.Param("shopId")
	err := ctrl.service.sendShop(shopID, req.Message)
	if err != nil {
		log.Println("Error occurred while send noti. to shop: ", err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{})
}

func (ctrl *NotificationSendController) HandleSendCustomerNotification(c *gin.Context) {
	var req NotificationSendRequest
	c.ShouldBindJSON(&req)
	customerID := c.Param("customerId")
	err := ctrl.service.sendCustomer(customerID, req.Message)
	if err != nil {
		log.Println("Error occurred while send noti. to customer: ", err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{})
}

func NewNotificationSendController(service NotificationService) *NotificationSendController {
	controller := new(NotificationSendController)
	controller.service = service
	return controller
}
