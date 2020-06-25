package main

import (
	log "github.com/micro/go-micro/v2/logger"
	"github.com/micro/go-micro/v2"
	"job/handler"
	"job/subscriber"

	job "job/proto/job"
)

func main() {
	// New Service
	service := micro.NewService(
		micro.Name("go.micro.service.job"),
		micro.Version("latest"),
	)

	// Initialise service
	service.Init()

	// Register Handler
	job.RegisterJobHandler(service.Server(), new(handler.Job))

	// Register Struct as Subscriber
	micro.RegisterSubscriber("go.micro.service.job", service.Server(), new(subscriber.Job))

	// Run service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
