package subscriber

import (
	"context"
	log "github.com/micro/go-micro/v2/logger"

	account "account/proto/account"
)

type Account struct{}

func (e *Account) Handle(ctx context.Context, msg *account.Message) error {
	log.Info("Handler Received message: ", msg.Say)
	return nil
}

func Handler(ctx context.Context, msg *account.Message) error {
	log.Info("Function Received message: ", msg.Say)
	return nil
}
