package main

import (
	"log"

	v1 "github.com/bells307/adv-service/internal/infrastructure/delivery/http/v1"
	"github.com/bells307/adv-service/internal/infrastructure/repository"
	"github.com/bells307/adv-service/internal/usecase"
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

	advRepo := repository.NewAdvertismentMongoDBRepository(mongoClient)
	advUsecase := usecase.NewAdvertismentUsecase(advRepo)
	advHandler := v1.NewAdvertismentHandler(advUsecase)
	advHandler.Register(router.Group("/api"))

	router.Run("localhost:10000")
}
