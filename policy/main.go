package main

import (
	log "github.com/micro/go-micro/v2/logger"
	"github.com/micro/go-micro/v2"
	"policy/handler"
	"policy/subscriber"

	policy "policy/proto/policy"
)

func main() {
	// New Service
	service := micro.NewService(
		micro.Name("go.micro.service.policy"),
		micro.Version("latest"),
	)

	// Initialise service
	service.Init()

	// Register Handler
	policy.RegisterPolicyHandler(service.Server(), new(handler.Policy))

	// Register Struct as Subscriber
	micro.RegisterSubscriber("go.micro.service.policy", service.Server(), new(subscriber.Policy))

	// Run service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
