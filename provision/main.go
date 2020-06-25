package main

import (
	log "github.com/micro/go-micro/v2/logger"
	"github.com/micro/go-micro/v2"
	"provision/handler"
	"provision/subscriber"

	provision "provision/proto/provision"
)

func main() {
	// New Service
	service := micro.NewService(
		micro.Name("go.micro.service.provision"),
		micro.Version("latest"),
	)

	// Initialise service
	service.Init()

	// Register Handler
	provision.RegisterProvisionHandler(service.Server(), new(handler.Provision))

	// Register Struct as Subscriber
	micro.RegisterSubscriber("go.micro.service.provision", service.Server(), new(subscriber.Provision))

	// Run service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
