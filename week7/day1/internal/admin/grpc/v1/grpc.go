package v1

import (
	admin_v1 "github.com/tuannguyenandpadcojp/go-training/week7/day1/internal/pb/v1"
)

type AdminService struct {
	admin_v1.UnimplementedAdminServiceServer
}
