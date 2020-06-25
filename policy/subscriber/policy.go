package subscriber

import (
	"context"
	log "github.com/micro/go-micro/v2/logger"

	policy "policy/proto/policy"
)

type Policy struct{}

func (e *Policy) Handle(ctx context.Context, msg *policy.Message) error {
	log.Info("Handler Received message: ", msg.Say)
	return nil
}

func Handler(ctx context.Context, msg *policy.Message) error {
	log.Info("Function Received message: ", msg.Say)
	return nil
}
