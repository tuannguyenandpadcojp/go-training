package usecase

import "context"

type CreateClientRequest struct{}
type CreateClientResponse struct{}

type ClientCreator interface {
	Create(ctx context.Context, req CreateClientRequest) (CreateClientResponse, error)
}
