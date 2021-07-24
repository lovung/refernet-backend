package usecase

import (
	"context"
	"refernet/internal/dto"
)

// UseCase interface to do an action
type UseCase interface {
	Do(ctx context.Context, input dto.Input) (output dto.Output, err error)
}
