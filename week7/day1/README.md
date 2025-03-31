# Tenant Management Application

This application provides a tenant management system with two gRPC services: **Admin Service** and **Client Service**. It is designed to serve both SaaS providers and their clients.

## Overview

- **Admin Service**:  
  This service is used by the SaaS provider's administrators to manage tenants, users, and configurations. It provides administrative functionalities such as creating tenants, managing users, and monitoring system usage.

- **Client Service**:
  This service is used by the end-users of the SaaS clients. It provides functionalities specific to the client's users, such as accessing tenant-specific resources and performing user-level operations.

## Code Structure

The project is organized as follows:
```
.
├── cmd # Entry point for the server
├── database
│   └── migrations # Database migration files
├── docs # Documentation files
├── internal
│   ├── admin # Admin service implementation
│   │   ├── grpc
│   │   │   └── v1
│   │   └── service # An implementation of the Admin usecases
│   ├── client # Client service implementation
│   │   ├── grpc
│   │   │   └── v1
│   │   └── service # An implementation of the Client usecases
│   ├── config # Configuration loader
│   ├── domain # Domain layer
│   │   └── repository # Repository interfaces that define the data access methods
│   ├── infrastructure # Infrastructure layer
│   ├── pb # Generated code from protobuf definitions
│   │   └── v1
│   └── usecase # Usecase layer
│       ├── admin # Usecases for the Admin service
│       └── client # Usecases for the Client service
├── proto # Protobuf definitions
│   └── v1
└── test # Test files
```

### Explanation of Key Components

1. **`/cmd/main.go`**
   The main entry point of the application. It initializes the gRPC server, registers the services, and handles graceful shutdown.

2. **`/internal/admin/grpc/v1/grpc.go`**
   Contains the implementation of the gRPC Admin Service. This service uses the `UnimplementedAdminServiceServer` from the generated protobuf code.

3. **`/internal/client/grpc/v1/grpc.go`**
   Contains the implementation of the gRPC Client Service. This service uses the `UnimplementedClientServiceServer` from the generated protobuf code.

4. **`/internal/config/config.go`**
   Handles the loading of configuration values from environment variables or configuration files.

5. **`/internal/pb/v1`**
   Contains the generated protobuf code for the gRPC services. This includes the service definitions and message types.

6. **`/internal/domain`**
   Contains the domain layer of the application, including repository interfaces that define the data access methods.

7. **`/internal/usecase`**
    Contains the use case layer of the application, which implements the business logic for both the Admin and Client services.

8. **`/internal/infrastructure`**
    Contains the infrastructure layer of the application, including database connections and other external service integrations.

9. **`/docs`**
   Contains documentation files, including setup instructions and API specifications.

## How to Run
1. **Setup application dependencies**
    Use the following command to install and run the dependencies:
    ```bash
    make up
    ```

1. **Set Up Configuration**
   Config environment variables for the application. You can create a `.local.env` file in the root directory by copying the `.env.example` file and modifying it according to your environment. The application will load the configuration from this file.

2. **Run the Application**
   Use the following command to start the server:
   ```bash
   make run
   ```

## Features

- **Admin Service**:
  - Manage tenants and users.
  - Monitor system usage.
  - Configure tenant-specific settings.

- **Client Service**:
  - Access tenant-specific resources.
  - Perform user-level operations.

## Future Enhancements

- Add authentication and authorization for both services.
- Implement logging and monitoring for better observability.
- Extend the Admin Service with advanced analytics and reporting.
