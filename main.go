package main

import (
	"github.com/SaiNageswarS/go-api-boot/server"
	pb "github.com/kotlang/localizationGo/generated"
	"github.com/rs/cors"
)

var grpcPort = ":50051"
var webPort = ":8081"

func main() {
	inject := NewInject()
	inject.CloudFns.LoadSecretsIntoEnv()

	corsConfig := cors.New(
		cors.Options{
			AllowedHeaders: []string{"*"},
		})

	bootServer := server.NewGoApiBoot(corsConfig)
	pb.RegisterLabelLocalizationServer(bootServer.GrpcServer, inject.LocalizationService)
	pb.RegisterLocalizationAdminServer(bootServer.GrpcServer, inject.LocalizationAdminService)

	bootServer.Start(grpcPort, webPort)
}
