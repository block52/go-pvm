# Go-PVM (Poker Virtual Machine)

A high-performance poker engine written in Go, ported from the TypeScript [Block52 Poker VM](https://github.com/block52/poker-vm).

## Overview

Go-PVM is a complete poker game engine that implements Texas Hold'em with support for multiple game formats:
- Cash games
- Sit & Go tournaments
- Multi-table tournaments

## Features

- **Texas Hold'em**: Full implementation of Texas Hold'em poker
- **Multiple Game Formats**: Support for cash games, sit-and-go, and tournaments
- **Action Validation**: Comprehensive validation of all player actions
- **Hand Evaluation**: Accurate poker hand ranking and winner determination
- **Betting Logic**: Support for blinds, antes, and rake
- **RPC Server**: HTTP/JSON-RPC interface for game interaction
- **Thread-Safe**: Designed for concurrent game handling

## Project Structure

```
go-pvm/
├── cmd/
│   └── server/              # RPC server entry point
├── internal/
│   ├── engine/              # Core poker engine
│   │   ├── actions/         # Poker actions (bet, call, raise, etc.)
│   │   ├── base/            # Base interfaces and implementations
│   │   ├── managers/        # Game state managers
│   │   └── holdem/          # Texas Hold'em implementation
│   ├── models/              # Data models (Player, Deck, etc.)
│   ├── types/               # Type definitions and interfaces
│   ├── utils/               # Utility functions
│   └── rpc/                 # RPC handler
├── tests/                   # Integration tests
└── APPROACH.md              # Porting strategy and implementation guide
```

## Quick Start

### Prerequisites

- Go 1.21 or higher

### Installation

```bash
git clone https://github.com/block52/go-pvm.git
cd go-pvm
go mod download
```

### Running the Server

```bash
go run cmd/server/main.go
```

The server will start on `http://localhost:8545`

### Running Tests

```bash
go test ./...
```

## Development

See [APPROACH.md](APPROACH.md) for the complete porting strategy and implementation details.

### .gitignore

The repository uses the existing Visual Studio .gitignore. Add the following Go-specific entries:

```
# Go-specific
*.exe
*.exe~
*.dll
*.so
*.dylib
*.test
*.out
go.work
vendor/
server
```

### Current Status

- [x] Project structure initialized
- [x] Core types and interfaces defined
- [x] Basic models (Player, Deck) implemented
- [ ] Managers implementation (in progress)
- [ ] Actions implementation (pending)
- [ ] Texas Hold'em game engine (pending)
- [ ] Hand evaluation (pending)
- [ ] RPC layer (pending)
- [ ] Full test suite (pending)

## API

The RPC server exposes the following endpoints:

- `GET /` - Server information
- `GET /health` - Health check
- `POST /` - JSON-RPC game commands (coming soon)

## License

MIT License - see [LICENSE](LICENSE) file for details

## Contributing

This is an active port from the TypeScript implementation. See APPROACH.md for the porting strategy and roadmap.

## Related Projects

- [poker-vm](https://github.com/block52/poker-vm) - Original TypeScript implementation
- [@block52/poker-vm-sdk](https://www.npmjs.com/package/@block52/poker-vm-sdk) - TypeScript SDK
