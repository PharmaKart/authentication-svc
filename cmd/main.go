package main

import (
	"net"

	"github.com/PharmaKart/authentication-svc/internal/handlers"
	pb "github.com/PharmaKart/authentication-svc/internal/proto"
	"github.com/PharmaKart/authentication-svc/internal/repositories"
	"github.com/PharmaKart/authentication-svc/pkg/config"
	"github.com/PharmaKart/authentication-svc/pkg/utils"

	"google.golang.org/grpc"
	"gorm.io/gorm"
)

func main() {
	// Load configuration
	cfg := config.LoadConfig()

	// Initialize logger
	utils.InitLogger()

	// Initialize repositories
	userRepo := repositories.NewUserRepository(&gorm.DB{})
	customerRepo := repositories.NewCustomerRepository(&gorm.DB{})

	// Initialize handlers
	authHandler := handlers.NewAuthHandler(userRepo, customerRepo, cfg.JWTSecret)

	// Initialize gRPC server
	lis, err := net.Listen("tcp", ":"+cfg.Port)

	if err != nil {
		utils.Logger.Fatal("Failed to listen", map[string]interface{}{
			"error": err,
		})
	}

	grpcServer := grpc.NewServer()
	pb.RegisterAuthServiceServer(grpcServer, authHandler)

	utils.Info("Starting authentication service", map[string]interface{}{
		"port": cfg.Port,
	})

	if err := grpcServer.Serve(lis); err != nil {
		utils.Logger.Fatal("Failed to serve", map[string]interface{}{
			"error": err,
		})
	}
}
