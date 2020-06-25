package main

import (
	log "github.com/micro/go-micro/v2/logger"
	"github.com/micro/go-micro/v2"
	"catalog/handler"
	"catalog/subscriber"

	catalog "catalog/proto/catalog"
)

func main() {
	// New Service
	service := micro.NewService(
		micro.Name("go.micro.service.catalog"),
		micro.Version("latest"),
	)

	// Initialise service
	service.Init()

	// Register Handler
	catalog.RegisterCatalogHandler(service.Server(), new(handler.Catalog))

	// Register Struct as Subscriber
	micro.RegisterSubscriber("go.micro.service.catalog", service.Server(), new(subscriber.Catalog))

	// Run service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
