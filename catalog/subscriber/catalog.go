package subscriber

import (
	"context"
	log "github.com/micro/go-micro/v2/logger"

	catalog "catalog/proto/catalog"
)

type Catalog struct{}

func (e *Catalog) Handle(ctx context.Context, msg *catalog.Message) error {
	log.Info("Handler Received message: ", msg.Say)
	return nil
}

func Handler(ctx context.Context, msg *catalog.Message) error {
	log.Info("Function Received message: ", msg.Say)
	return nil
}
