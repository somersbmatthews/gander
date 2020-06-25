package main

import (
	log "github.com/micro/go-micro/v2/logger"
	"github.com/micro/go-micro/v2"
	"event/handler"
	"event/subscriber"

	event "event/proto/event"
)

func main() {
	// New Service
	service := micro.NewService(
		micro.Name("go.micro.service.event"),
		micro.Version("latest"),
	)

	// Initialise service
	service.Init()

	// Register Handler
	event.RegisterEventHandler(service.Server(), new(handler.Event))

	// Register Struct as Subscriber
	micro.RegisterSubscriber("go.micro.service.event", service.Server(), new(subscriber.Event))

	// Run service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
