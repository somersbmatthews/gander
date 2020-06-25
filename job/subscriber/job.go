package subscriber

import (
	"context"
	log "github.com/micro/go-micro/v2/logger"

	job "job/proto/job"
)

type Job struct{}

func (e *Job) Handle(ctx context.Context, msg *job.Message) error {
	log.Info("Handler Received message: ", msg.Say)
	return nil
}

func Handler(ctx context.Context, msg *job.Message) error {
	log.Info("Function Received message: ", msg.Say)
	return nil
}
