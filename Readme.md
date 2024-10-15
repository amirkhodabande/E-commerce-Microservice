# E-commerce Microservice

This project is a microservice-based architecture for an e-commerce platform. It is designed to be a simple, scalable, and maintainable solution.
Each service is a separate entity that can be deployed and scaled independently, following Domain-Driven Design (DDD) principles.

## Services

### User Service

The User Service is responsible for managing user data, authentication, and authorization.

### Clients
Each service provides clients that the gateway uses to communicate with them. This approach ensures transparency and ease of access to the services.
Additionally, a validation layer is included for these clients.

**structure**
```json
clients
├── user
│   ├── http
│   │   ├── clients
│   │   │   ├── example_client.go
│   │   └── data_objects
│   │       ├── example_data_object.go
│   │── grpc/...
│   │   ├── clients
│   │   │   ├── example_client.go
│   │   └── data_objects
│   │       ├── example_data_object.go
├── antoher service
│   ├── http
│   │   ├── clients
│   │   │   ├── example_client.go
│   │   └── data_objects
│   │       ├── example_data_object.go
│   │── grpc/...
│   │   ├── clients
│   │   │   ├── example_client.go
│   │   └── data_objects
│   │       ├── example_data_object.go
```

## Gateway

The Gateway is responsible for routing requests to the appropriate service. It is designed to be a simple and scalable solution for an e-commerce platform.
*It is the only service that is exposed to the outside world.*

*Will use clients of each service*

## Inprogress...





