# Go Domain-Driven Project Boilerplate

A production-ready Go application boilerplate organized around business domains, following clean architecture and domain-driven design principles with dependency injection using Wire.

## Project Structure

```
.
├── cmd/                  # Application entry points
│   ├── main.go           # Main application
│   ├── wire.go           # Dependency injection wire setup
│   └── wire_gen.go       # Generated wire code
├── internal/             # Private application code
│   ├── api/              # API definitions and interfaces
│   ├── application/      # Application services and use cases
│   │   ├── container.go  # Application service container
│   │   ├── controller.go # Application service controller
│   │   ├── errors.go     # Domain-specific errors
│   │   ├── model.go      # Domain models
│   │   ├── repository.go # Repository interface and implementation
│   │   └── wire_set.go   # Wire dependency set
│   ├── auth/             # Authentication mechanisms
│   ├── config/           # Configuration management
│   ├── connection/       # External service connections
│   ├── database/         # Database connections and migrations
│   └── user/             # User domain
│       ├── container.go  # User service container
│       ├── controller.go # User controller
│       ├── errors.go     # User-specific errors
│       ├── model.go      # User domain models
│       ├── repository.go # User repository interface and implementation
│       └── wire_set.go   # Wire dependency set
├── util/                 # Utility packages
│   ├── ctx/              # Context utilities
│   └── http/             # HTTP utilities
├── .env                  # Environment variables configuration
├── go.mod                # Go module definition
├── go.sum                # Go module checksums
└── Makefile              # Build and development commands
```

## Features

- **Domain-Driven Design**: Business logic organized around domains
- **Clean Architecture**: Clear separation of concerns between domains
- **Dependency Injection**: Using Google's Wire for DI
- **Environment Configuration**: Using .env for configuration
- **Repository Pattern**: Clean data access abstractions
- **Structured Error Handling**: Domain-specific error types
- **HTTP Utilities**: Common HTTP patterns and middleware
- **Context Management**: Specialized context utilities
- **Makefile Support**: Simplified command to run database migrations

## Getting Started

### Prerequisites

- Go 1.18 or higher

### Setup

1. Clone the repository:

   ```bash
   git clone https://github.com/yourusername/your-project-name.git
   cd your-project-name
   ```

2. Install dependencies:

   ```bash
   go mod download
   ```

3. Set up environment variables:

   ```bash
   # Ensure .env file is configured properly
   ```

4. Run the application:
   ```bash
   go run cmd/
   ```

## Development

### Domain Structure Pattern

Each domain follows this structure:

1. **model.go**: Domain entities and value objects
2. **repository.go**: Data access interfaces and implementations
3. **controller.go**: Business logic and use cases
4. **container.go**: Service container for dependency management
5. **errors.go**: Domain-specific error types
6. **wire_set.go**: Wire dependency injection set

### Adding a New Domain

1. Create a new directory in `internal/` for your domain
2. Create the 6 files mentioned above:
   - Define domain models in `model.go`
   - Define repository interface in `repository.go`
   - Implement business logic in `controller.go`
   - Set up dependencies in `container.go`
   - Define domain errors in `errors.go`
   - Configure wire set in `wire_set.go`
3. Add the domain's wire set to `cmd/wire.go`
4. Run `wire` to generate the updated dependency graph

### Common Development Tasks

```bash
# Generate Wire dependency injection code
go run github.com/google/wire/cmd/wire

```

## License

This project is licensed under the MIT License - see the LICENSE file for details.
