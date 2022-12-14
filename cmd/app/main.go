package main

import (
	"log"

	adv_repository "github.com/bells307/adv-service/internal/domain/advertisment/repository"
	adv_service "github.com/bells307/adv-service/internal/domain/advertisment/service"
	adv_handler "github.com/bells307/adv-service/internal/handler/advertisment/v1"
	"github.com/bells307/adv-service/pkg/mongodb"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	mongoConfig := mongodb.MongoDBConfig{
		Uri:    "mongodb://admin:admin@localhost:27017",
		DbName: "adv-service",
	}

	mongoClient, err := mongodb.NewMongoDB(mongoConfig)
	if err != nil {
		log.Fatalf("can't connect to MongoDB: %v", err)
	}

	advRepo := adv_repository.NewAdvertismentMongoDBRepository(mongoClient)
	advService := adv_service.NewAdvertismentService(advRepo)
	advHandler := adv_handler.NewAdvertismentHandler(advService)
	advHandler.Register(router.Group("/api"))

	router.Run("localhost:10000")
}
