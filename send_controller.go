package main

import "github.com/gin-gonic/gin"

import "net/http"

import "log"

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
		log.Fatalln("Error occurred while send noti. to shop: ", err.Error())
		c.JSON(http.StatusInternalServerError, err.Error())
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
		log.Fatalln("Error occurred while send noti. to customer: ", err.Error())
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, gin.H{})
}

func NewNotificationSendController(service NotificationService) *NotificationSendController {
	controller := new(NotificationSendController)
	controller.service = service
	return controller
}
