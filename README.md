# Sapphire ICICI

A Go-based web application built with the Gin framework for ICICI Bank operations.

## Prerequisites

- Go 1.19 or higher
- Git

## Installation

### 1. Clone the repository
```bash
git clone <your-repo-url>
cd SapphireICICI
```

### 2. Setup the project

**Option A: Using setup script (recommended)**
```bash
./setup.sh
```

**Option B: Manual setup**
```bash
# Download dependencies
go mod download

# Install development tools
go install github.com/evilmartians/lefthook@latest
go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest

# Setup git hooks
lefthook install
```

## Usage

### Running the application
```bash
go run ./cmd/server
```

### Building the application
```bash
go build -o tmp/server ./cmd/server
```

### Running tests
```bash
go test ./...
```

### Code formatting and linting
```bash
# Format code
go fmt ./...

# Run linter
golangci-lint run

# Run all pre-commit checks
lefthook run pre-commit
```

## Development

### Git Hooks

This project uses [lefthook](https://github.com/evilmartians/lefthook) for Git hooks automation:

**Pre-commit hooks:**
- `go fmt` - Code formatting
- `go vet` - Static analysis
- `go test` - Run tests
- `golangci-lint` - Linting

**Pre-push hooks:**
- `go build` - Build verification

### Project Structure
```
SapphireICICI/
├── cmd/
│   └── server/
│       └── main.go          # Application entry point
├── internal/
│   ├── config/              # Configuration management
│   ├── env/                 # Environment variables
│   ├── errors/              # Error handling
│   ├── logger/              # Logging utilities
│   ├── middleware/          # Gin middleware
│   └── utils/               # Utility functions
├── logs/                    # Log files
├── tmp/                     # Build artifacts
├── go.mod                   # Go module file
├── go.sum                   # Dependency checksums
├── lefthook.yml             # Git hooks configuration
└── README.md
```

## Go vs npm Commands

| npm | Go | Purpose |
|-----|-----|---------|
| `npm install` | `go mod download` | Download dependencies |
| `npm install --save` | `go get <package>` | Add new dependency |
| `npm update` | `go get -u ./...` | Update dependencies |
| `npm run build` | `go build` | Build application |
| `npm start` | `go run ./cmd/server` | Run application |
| `npm test` | `go test ./...` | Run tests |

## Key Features of Go Dependency Management

- **No `node_modules`** - Go downloads dependencies to a global cache
- **No `package-lock.json`** - Go uses `go.sum` for version locking
- **Automatic dependency resolution** - Dependencies are downloaded when running `go run` or `go build`
- **Built-in tooling** - No need for separate package managers

## Quick Start

For new team members, simply run:
```bash
git clone <your-repo-url>
cd SapphireICICI
go run ./cmd/server  # This will automatically download dependencies
```

## Contributing

1. Fork the repository
2. Create a feature branch: `git checkout -b feature/your-feature`
3. Commit your changes: `git commit -m "feat: add your feature"`
4. Push to the branch: `git push origin feature/your-feature`
5. Submit a pull request
