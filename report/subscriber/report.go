package subscriber

import (
	"context"
	log "github.com/micro/go-micro/v2/logger"

	report "report/proto/report"
)

type Report struct{}

func (e *Report) Handle(ctx context.Context, msg *report.Message) error {
	log.Info("Handler Received message: ", msg.Say)
	return nil
}

func Handler(ctx context.Context, msg *report.Message) error {
	log.Info("Function Received message: ", msg.Say)
	return nil
}
