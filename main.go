package main

import (
	"os"

	"github.com/SaiNageswarS/go-api-boot/logger"
	"github.com/SaiNageswarS/go-api-boot/server"
	"github.com/joho/godotenv"
	pb "github.com/kotlang/localizationGo/generated"
	"go.uber.org/zap"
)

var grpcPort = ":50051"
var webPort = ":8081"

func init() {
	err := godotenv.Load()
	if err != nil {
		logger.Error("Error loading .env file", zap.Error(err))
	}
}

func main() {
	// go-api-boot picks up keyvault name from environment variable.
	os.Setenv("AZURE-KEYVAULT-NAME", "kotlang-secrets")
	server.LoadSecretsIntoEnv(true)
	inject := NewInject()

	bootServer := server.NewGoApiBoot()
	pb.RegisterLabelLocalizationServer(bootServer.GrpcServer, inject.LocalizationService)
	pb.RegisterLocalizationAdminServer(bootServer.GrpcServer, inject.LocalizationAdminService)

	bootServer.Start(grpcPort, webPort)
}
