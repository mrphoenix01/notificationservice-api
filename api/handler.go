// api/handler.go
package api

import (
	"context"
	"net/http"
	"notification_service/dataservice"
	"notification_service/kafka"
	"notification_service/model"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
)

// SubscribeHandler - Handles subscription requests
func SubscribeHandler(c *gin.Context) {
	var subscription model.Subscription
	if err := c.BindJSON(&subscription); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	collection := dataservice.GetCollection("notification_service", "subscriptions")
	_, err := collection.InsertOne(context.TODO(), subscription)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to subscribe"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Subscription successful"})
}

// SendNotificationHandler - Sends a notification to all subscribers of a specific topic
func SendNotificationHandler(c *gin.Context) {
	var notification model.Notification
	if err := c.BindJSON(&notification); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Send notification message to Kafka topic
	if err := kafka.PublishMessage(notification.Topic, notification); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to send notification"})
		return
	}

	collection := dataservice.GetCollection("notification_service", "notifications")
	_, err := collection.InsertOne(context.TODO(), notification)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save notification"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Notification sent"})
}

// UnsubscribeHandler - Unsubscribes a user from specified topics
func UnsubscribeHandler(c *gin.Context) {
	var unsubscribeRequest struct {
		UserID string   `json:"user_id"`
		Topics []string `json:"topics"`
	}
	if err := c.BindJSON(&unsubscribeRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	collection := dataservice.GetCollection("notification_service", "subscriptions")
	filter := bson.M{"user_id": unsubscribeRequest.UserID}
	update := bson.M{"$pull": bson.M{"topics": bson.M{"$in": unsubscribeRequest.Topics}}}

	_, err := collection.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to unsubscribe"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Unsubscribed successfully"})
}

// GetSubscriptionsHandler - Fetches user subscriptions
func GetSubscriptionsHandler(c *gin.Context) {
	userID := c.Param("user_id")

	collection := dataservice.GetCollection("notification_service", "subscriptions")
	var subscriptions []model.Subscription
	cursor, err := collection.Find(context.TODO(), bson.M{"user_id": userID})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve subscriptions"})
		return
	}
	defer cursor.Close(context.TODO())

	for cursor.Next(context.TODO()) {
		var subscription model.Subscription
		if err := cursor.Decode(&subscription); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error decoding subscription"})
			return
		}
		subscriptions = append(subscriptions, subscription)
	}

	c.JSON(http.StatusOK, gin.H{"user_id": userID, "subscriptions": subscriptions})
}
