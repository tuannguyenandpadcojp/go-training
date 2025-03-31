# Proto definition for gRPC services

This document outlines the steps to define a gRPC service using Protocol Buffers (Proto).
Generate code using Buf.

## Prerequisites
- Install the Buf CLI tool: https://docs.buf.build/installation
- Install Go plugins for the protocol compiler: https://grpc.io/docs/languages/go/quickstart/#prerequisites

## Protocol Buffers (Proto) Style Guide
- [proto 3](https://protobuf.dev/programming-guides/proto3/)
- [Buf Style Guide](https://docs.buf.build/style-guide) 

## Get Started
1. Create a new directory for your proto files
```bash
mkdir -p proto/v1
```

2. Create new proto files
- Create a proto file for the Client service that includes the service definition and request/response messages
```bash
touch proto/v1/admin_service.proto
```

- Create a proto file for the Client service
```bash
touch proto/v1/client_service.proto
```

- Create a proto file for the data model that includes the data model definition
```bash
touch proto/v1/client.proto
```

3. Using Buf to generate code
- Create a new buf.yaml file to configure the Buf build
```bash
buf config init
```

- Create a new buf.gen.yaml file to configure the Buf generation
```bash
touch buf.gen.yaml
```

- Generate code using Buf
```bash
buf generate
```

4. Verify the generated code
You will find the generated code in the `proto/v1` directory.
