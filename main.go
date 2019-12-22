package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	repository := NewRedisRepository("localhost:6379", "", 0)
	apiWrapper := NewFcmApiWrapper("")
	service := NewNotificationService(apiWrapper, repository)
	subscriptionCtrl := NewSubscriptionController(*service)
	sendCtrl := NewNotificationSendController(*service)
	r.POST("/shops/:shopId/subscribe", subscriptionCtrl.HandleShopSubscribe)
	r.POST("/customers/:customerId/subscribe", subscriptionCtrl.HandleCustomerSubscribe)
	r.POST("/shops/:shopId/send", sendCtrl.HandleSendShopNotification)
	r.POST("/cusotmers/:customerId/send", sendCtrl.HandleSendCustomerNotification)
	r.Run("0.0.0.0:8000")
}
