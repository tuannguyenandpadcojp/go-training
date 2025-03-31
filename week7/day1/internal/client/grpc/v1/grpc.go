package v1

import (
	client_v1 "github.com/tuannguyenandpadcojp/go-training/week7/day1/internal/pb/v1"
)

type ClientService struct {
	client_v1.UnimplementedClientServiceServer
}
