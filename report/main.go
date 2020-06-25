package main

import (
	log "github.com/micro/go-micro/v2/logger"
	"github.com/micro/go-micro/v2"
	"report/handler"
	"report/subscriber"

	report "report/proto/report"
)

func main() {
	// New Service
	service := micro.NewService(
		micro.Name("go.micro.service.report"),
		micro.Version("latest"),
	)

	// Initialise service
	service.Init()

	// Register Handler
	report.RegisterReportHandler(service.Server(), new(handler.Report))

	// Register Struct as Subscriber
	micro.RegisterSubscriber("go.micro.service.report", service.Server(), new(subscriber.Report))

	// Run service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
