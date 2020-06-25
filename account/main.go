package main

import (
	log "github.com/micro/go-micro/v2/logger"
	"github.com/micro/go-micro/v2"
	"account/handler"
	"account/subscriber"

	account "account/proto/account"
)

func main() {
	// New Service
	service := micro.NewService(
		micro.Name("go.micro.service.account"),
		micro.Version("latest"),
	)

	// Initialise service
	service.Init()

	// Register Handler
	account.RegisterAccountHandler(service.Server(), new(handler.Account))

	// Register Struct as Subscriber
	micro.RegisterSubscriber("go.micro.service.account", service.Server(), new(subscriber.Account))

	// Run service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
