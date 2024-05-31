package telegram

import "context"

//go:generate ./../../../bin/mockery --name Usecase
type Usecase interface {
	SendMessage(ctx context.Context, chatID, message string) error
}
