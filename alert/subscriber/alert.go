package subscriber

import (
	"context"
	log "github.com/micro/go-micro/v2/logger"

	alert "alert/proto/alert"
)

type Alert struct{}

func (e *Alert) Handle(ctx context.Context, msg *alert.Message) error {
	log.Info("Handler Received message: ", msg.Say)
	return nil
}

func Handler(ctx context.Context, msg *alert.Message) error {
	log.Info("Function Received message: ", msg.Say)
	return nil
}
