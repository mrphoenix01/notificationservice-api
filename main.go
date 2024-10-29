package main

import (
	"log"
	"notification_service/api"
	"notification_service/dataservice"
	"notification_service/kafka"

	"github.com/gin-gonic/gin"
)

func main() {

	dataservice.ConnectMongoDB()
	defer func() {
		if err := dataservice.Client.Disconnect(nil); err != nil {
			log.Fatal(err)
		}
	}()

	kafka.InitializeKafkaProducer([]string{"localhost:9092"})

	router := gin.Default()
	api.SetupRoutes(router)

	log.Println("Starting server on :8080")
	if err := router.Run(":8082"); err != nil {
		log.Fatal("Server error:", err)
	}
}
