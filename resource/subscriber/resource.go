package subscriber

import (
	"context"
	log "github.com/micro/go-micro/v2/logger"

	resource "resource/proto/resource"
)

type Resource struct{}

func (e *Resource) Handle(ctx context.Context, msg *resource.Message) error {
	log.Info("Handler Received message: ", msg.Say)
	return nil
}

func Handler(ctx context.Context, msg *resource.Message) error {
	log.Info("Function Received message: ", msg.Say)
	return nil
}
