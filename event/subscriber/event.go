package subscriber

import (
	"context"
	log "github.com/micro/go-micro/v2/logger"

	event "event/proto/event"
)

type Event struct{}

func (e *Event) Handle(ctx context.Context, msg *event.Message) error {
	log.Info("Handler Received message: ", msg.Say)
	return nil
}

func Handler(ctx context.Context, msg *event.Message) error {
	log.Info("Function Received message: ", msg.Say)
	return nil
}
