// main.go
package main

import (
	"log"
	"notification_service/api"
	"notification_service/dataservice"
	"notification_service/kafka" // Import the Kafka package

	"github.com/gin-gonic/gin"
)

func main() {
	// Connect to MongoDB
	dataservice.ConnectMongoDB()
	defer func() {
		if err := dataservice.Client.Disconnect(nil); err != nil {
			log.Fatal(err)
		}
	}()

	// Initialize Kafka Producer
	kafka.InitializeKafkaProducer([]string{"localhost:9092"}) // Adjust to match your Kafka broker(s)

	// Initialize Gin router
	router := gin.Default()
	api.SetupRoutes(router)

	// Start the server
	log.Println("Starting server on :8080")
	if err := router.Run(":8081"); err != nil {
		log.Fatal("Server error:", err)
	}
}
