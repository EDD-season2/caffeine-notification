package main

import (
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	redisAddr := os.Getenv("REDIS_ADDR")
	authHeader := os.Getenv("AUTH_HEADER")
	repository := NewRedisRepository(redisAddr, "", 0)
	apiWrapper := NewFcmApiWrapper(authHeader)
	service := NewNotificationService(apiWrapper, repository)
	subscriptionCtrl := NewSubscriptionController(*service)
	sendCtrl := NewNotificationSendController(*service)
	r.POST("/shops/:shopId/subscribe", subscriptionCtrl.HandleShopSubscribe)
	r.POST("/customers/:customerId/subscribe", subscriptionCtrl.HandleCustomerSubscribe)
	r.POST("/shops/:shopId/send", sendCtrl.HandleSendShopNotification)
	r.POST("/customers/:customerId/send", sendCtrl.HandleSendCustomerNotification)
	r.Run("0.0.0.0:8000")
}
