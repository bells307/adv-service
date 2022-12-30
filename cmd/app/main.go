package main

import (
	advertismentgrpc "github.com/bells307/adv-service/internal/infrastructure/delivery/grpc/advertisment"
	"google.golang.org/grpc"
	"log"
	"net"

	"github.com/bells307/adv-service/cmd/app/config"
	"github.com/bells307/adv-service/internal/adapter/repository"
	v1 "github.com/bells307/adv-service/internal/infrastructure/delivery/http/v1"
	"github.com/bells307/adv-service/pkg/mongodb"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title	adv-service API
// @verion	0.1
func main() {
	cfg, err := config.LoadConfig(".env")
	if err != nil {
		log.Fatalf("error while loading application configuration: %v", err)
	}

	router := gin.Default()
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	mongoClient, err := mongodb.NewMongoDB(cfg.MongoDB)
	if err != nil {
		log.Fatalf("can't connect to MongoDB: %v", err)
	}

	advRepo := repository.NewAdvertismentMongoDBRepository(mongoClient)
	catRepo := repository.NewCategoryMongoDBRepository(mongoClient)

	advHandler := v1.NewAdvertismentHandler(advRepo, catRepo)
	advHandler.Register(router.Group("/api"))

	catHandler := v1.NewCategoryHandler(advRepo, catRepo)
	catHandler.Register(router.Group("/api"))

	go func() { router.Run(cfg.HttpListen) }()

	grpcServer := grpc.NewServer()
	advGrpc := advertismentgrpc.NewAdvertismentHandler(advRepo, catRepo)
	advGrpc.Register(grpcServer)
	listener, err := net.Listen("tcp", cfg.GrpcListen)
	if err != nil {
		log.Fatalf("can't create grpc listener: %v", err)
	}
	err = grpcServer.Serve(listener)
	if err != nil {
		log.Fatalf("error while starting grpc serving: %v", err)
	}
}
