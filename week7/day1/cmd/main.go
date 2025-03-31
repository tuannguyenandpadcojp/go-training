package main

import (
	"fmt"
	"net"
	"os"
	"os/signal"
	"syscall"

	_ "github.com/go-sql-driver/mysql" // MySQL driver
	"github.com/rs/zerolog"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	admin_grpc_v1 "github.com/tuannguyenandpadcojp/go-training/week7/day1/internal/admin/grpc/v1"
	client_grpc_v1 "github.com/tuannguyenandpadcojp/go-training/week7/day1/internal/client/grpc/v1"
	"github.com/tuannguyenandpadcojp/go-training/week7/day1/internal/config"
	pb_v1 "github.com/tuannguyenandpadcojp/go-training/week7/day1/internal/pb/v1"
)

func main() {
	fmt.Printf("config path: %s\n", os.Getenv("CONFIG_PATH"))
	cfg, err := config.LoadConfig(os.Getenv("CONFIG_PATH"))
	if err != nil {
		panic(err)
	}
	// Initialize logger
	logger := newLogger(cfg)

	// Initialize gRPC server
	server := newGRPCServer(cfg)
	server.RegisterService(&pb_v1.AdminService_ServiceDesc, &admin_grpc_v1.AdminService{})
	server.RegisterService(&pb_v1.ClientService_ServiceDesc, &client_grpc_v1.ClientService{})

	var grpcWaiter chan struct{}
	go func() {
		defer func() { grpcWaiter <- struct{}{} }()
		lis, err := net.Listen("tcp", cfg.GRPCAddr)
		if err != nil {
			logger.Err(err).Msgf("gRPC.server: failed to listen on address: %s", cfg.GRPCAddr)
			return
		}
		logger.Info().Msgf("gRPC.server: listening on address: %s", cfg.GRPCAddr)
		if err := server.Serve(lis); err != nil && err != grpc.ErrServerStopped {
			logger.Err(err).Msgf("gRPC.server: failed to serve on address: %s", cfg.GRPCAddr)
			return
		}
	}()

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	select {
	case <-c:
		logger.Info().Msg("Received shutdown signal")
	case <-grpcWaiter:
		logger.Info().Msg("gRPC.server: stopped")
	}
	// Graceful shutdown
	server.GracefulStop()
}

func newLogger(cfg *config.Config) *zerolog.Logger {
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	zerolog.SetGlobalLevel(zerolog.InfoLevel)
	if cfg.Debug {
		zerolog.SetGlobalLevel(zerolog.DebugLevel)
	}
	logger := zerolog.New(os.Stdout).With().Timestamp().Logger()
	return &logger
}

func newGRPCServer(cfg *config.Config) *grpc.Server {
	// Initialize gRPC server
	grpcServer := grpc.NewServer()
	if cfg.Env == "local" {
		reflection.Register(grpcServer)
	}
	return grpcServer
}
