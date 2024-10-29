// api/routes.go
package api

import (
	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.Engine) {
	router.POST("/subscribe", SubscribeHandler)
	router.POST("/notifications/send", SendNotificationHandler)
	router.POST("/unsubscribe", UnsubscribeHandler)
	router.GET("/subscriptions/:user_id", GetSubscriptionsHandler)
}
