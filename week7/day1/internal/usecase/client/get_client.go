package usecase

import (
	"context"

	"github.com/tuannguyenandpadcojp/go-training/week7/day1/internal/domain"
)

type GetClientRequest struct {
	UserAttributes domain.UserAttributes
}

type GetClientResponse struct {
	Client domain.Client
}

type IGetClient interface {
	GetClient(ctx context.Context, req GetClientRequest) (GetClientResponse, error)
}
