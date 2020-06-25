package main

import (
	log "github.com/micro/go-micro/v2/logger"
	"github.com/micro/go-micro/v2"
	"alert/handler"
	"alert/subscriber"

	alert "alert/proto/alert"
)

func main() {
	// New Service
	service := micro.NewService(
		micro.Name("go.micro.service.alert"),
		micro.Version("latest"),
	)

	// Initialise service
	service.Init()

	// Register Handler
	alert.RegisterAlertHandler(service.Server(), new(handler.Alert))

	// Register Struct as Subscriber
	micro.RegisterSubscriber("go.micro.service.alert", service.Server(), new(subscriber.Alert))

	// Run service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
