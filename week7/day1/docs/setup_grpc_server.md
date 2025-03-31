# Setting Up a gRPC Server

This document outlines the steps to initialize and configure a gRPC server as implemented in `cmd/main.go`.

## Summary

By following these steps, the gRPC server is set up to handle requests, log activities, and gracefully shut down when required. This implementation ensures robust and maintainable server behavior.

## Steps to Initialize the gRPC Server

1. **Load Configuration**
   Use the `config.LoadConfig` function to load the server configuration from the environment variable `CONFIG_PATH`. This configuration includes the gRPC address, logging level, and other runtime options.

   ```go
   cfg, err := config.LoadConfig(os.Getenv("CONFIG_PATH"))
   if err != nil {
       panic(err)
   }
   ```

2. **Initialize Logging**
   Create a logger using the `zerolog` library. The logging level is set based on the configuration (`Debug` mode or default `Info` level).

   ```go
   logger := newLogger(cfg)
   ```

3. **Create a New gRPC Server**
   Instantiate a gRPC server using `grpc.NewServer()`. If the environment is `local`, enable reflection for easier debugging and development.

   ```go
   server := newGRPCServer(cfg)
   ```

4. **Register gRPC Services**
   Add the required gRPC services to the server. In this case, the `AdminService` and `ClientService` are registered.

   ```go
   server.RegisterService(&pb_v1.AdminService_ServiceDesc, &admin_grpc_v1.AdminService{})
   server.RegisterService(&pb_v1.ClientService_ServiceDesc, &client_grpc_v1.ClientService{})
   ```

   - **AdminService**: A simple service that uses `UnimplementedAdminServiceServer` from the generated protobuf code.
   - **ClientService**: A simple service that uses `UnimplementedClientServiceServer` from the generated protobuf code.

5. **Start the gRPC Server**
   Start the gRPC server on the configured address (`cfg.GRPCAddr`). Use a goroutine to listen for incoming connections and handle errors.

   ```go
   go func() {
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
   ```

6. **Handle Shutdown Signals**
   Use the `os/signal` package to listen for termination signals (`os.Interrupt` or `syscall.SIGTERM`). When a signal is received, gracefully stop the gRPC server.

   ```go
   c := make(chan os.Signal, 1)
   signal.Notify(c, os.Interrupt, syscall.SIGTERM)
   select {
   case <-c:
       logger.Info().Msg("Received shutdown signal")
   case <-grpcWaiter:
       logger.Info().Msg("gRPC.server: stopped")
   }
   server.GracefulStop()
   ```

## Verification
To verify that the gRPC server is running correctly, you can use a gRPC client as Postman to send requests to the server. Ensure that the server is reachable at the configured address and that the services are responding as expected.
