package main

import (
	log "github.com/micro/go-micro/v2/logger"
	"github.com/micro/go-micro/v2"
	"resource/handler"
	"resource/subscriber"

	resource "resource/proto/resource"
)

func main() {
	// New Service
	service := micro.NewService(
		micro.Name("go.micro.service.resource"),
		micro.Version("latest"),
	)

	// Initialise service
	service.Init()

	// Register Handler
	resource.RegisterResourceHandler(service.Server(), new(handler.Resource))

	// Register Struct as Subscriber
	micro.RegisterSubscriber("go.micro.service.resource", service.Server(), new(subscriber.Resource))

	// Run service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
