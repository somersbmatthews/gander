package subscriber

import (
	"context"
	log "github.com/micro/go-micro/v2/logger"

	provision "provision/proto/provision"
)

type Provision struct{}

func (e *Provision) Handle(ctx context.Context, msg *provision.Message) error {
	log.Info("Handler Received message: ", msg.Say)
	return nil
}

func Handler(ctx context.Context, msg *provision.Message) error {
	log.Info("Function Received message: ", msg.Say)
	return nil
}
