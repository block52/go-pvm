# Go-PVM Porting Approach

This document outlines the strategy for porting the TypeScript Poker VM to Go.

## Source Analysis

The TypeScript poker engine is located at `/Users/lucascullen/Github/block52/poker-vm/pvm/ts/src` and consists of:

### Core Components

1. **Engine** (`src/engine/`)
   - `texasHoldem.ts` (75KB) - Main Texas Hold'em game engine
   - `actions/` - 20+ action implementations (fold, call, raise, bet, check, etc.)
   - `managers/` - Game state managers (dealer, blinds, payout, status, bet)
   - `base/` - Abstract base classes for poker game implementations
   - `types.ts` - Core interfaces (IPoker, IAction, IDealer, etc.)

2. **Models** (`src/models/`)
   - `player.ts` - Player model
   - `deck.ts` - Deck management with card dealing
   - `interfaces.ts` - Model interfaces

3. **Utils** (`src/utils/`)
   - `crypto.ts` - Cryptographic functions
   - `parsers.ts` - Data parsing utilities
   - `logger/` - Logging functionality

4. **RPC Layer** (`src/rpc.ts`, `src/index.ts`)
   - Express-based HTTP server
   - JSON-RPC handler
   - WebSocket support (optional)

5. **Testing**
   - 30+ comprehensive test files
   - Tests for each action type
   - Integration tests for full game scenarios
   - Tests for edge cases (heads-up, sit-and-go, rake, etc.)

### Key Architecture Patterns

1. **Interface-Driven Design**
   - `IPoker` - Core game operations
   - `IAction` - Action verification and execution
   - `IDealer` - Dealer position management
   - `IDealerPositionManager` - Position tracking

2. **Manager Pattern**
   - `BetManager` - Betting logic and validation
   - `DealerManager` - Dealer position rotation
   - `BlindsManager` - Blind posting logic
   - `PayoutManager` - Winner calculation and chip distribution
   - `StatusManager` - Player status transitions

3. **Action Pattern**
   - Each action (Bet, Call, Raise, Fold, etc.) is a separate class
   - Actions implement `verify()` and `execute()` methods
   - Actions are responsible for validating player state and updating game state

4. **Game Format Support**
   - Cash games
   - Sit & Go tournaments
   - Multi-table tournaments (via game options)

## Go Project Structure

```
go-pvm/
├── cmd/
│   └── server/              # RPC server entry point
│       └── main.go
├── internal/
│   ├── engine/              # Core poker engine
│   │   ├── base/            # Base interfaces and abstract implementations
│   │   ├── actions/         # Action implementations
│   │   │   ├── fold.go
│   │   │   ├── call.go
│   │   │   ├── raise.go
│   │   │   ├── bet.go
│   │   │   ├── check.go
│   │   │   ├── allin.go
│   │   │   └── ...
│   │   ├── managers/        # Game state managers
│   │   │   ├── dealer.go
│   │   │   ├── blinds.go
│   │   │   ├── payout.go
│   │   │   ├── status.go
│   │   │   └── bet.go
│   │   └── holdem/          # Texas Hold'em implementation
│   │       └── game.go
│   ├── models/              # Data models
│   │   ├── player.go
│   │   ├── deck.go
│   │   └── card.go
│   ├── types/               # Types and interfaces
│   │   ├── types.go
│   │   └── interfaces.go
│   ├── utils/               # Utility functions
│   │   ├── crypto.go
│   │   └── parsers.go
│   └── rpc/                 # RPC handler
│       └── handler.go
├── pkg/                     # Public packages (if needed)
├── tests/                   # Integration tests
├── go.mod
├── go.sum
├── APPROACH.md
└── README.md
```

## Porting Strategy

### Phase 1: Foundation (Week 1)
**Goal**: Establish core types, interfaces, and basic models

- [x] Initialize Go module
- [x] Create directory structure
- [x] Define core types (PlayerActionType, GameFormat, etc.)
- [x] Define core interfaces (IPoker, IAction, IDealer)
- [x] Implement basic models (Player, Deck, Card)
- [ ] Port utility functions (crypto, parsers)
- [ ] Set up testing framework

### Phase 2: Core Engine (Week 2-3)
**Goal**: Implement the poker game engine base and managers

- [ ] Implement PokerGameBase (abstract base class)
- [ ] Implement Deck with proper shuffling and dealing
- [ ] Implement DealerPositionManager
- [ ] Implement BlindsManager
- [ ] Implement BetManager
- [ ] Implement PayoutManager
- [ ] Implement StatusManager
- [ ] Write unit tests for each manager

### Phase 3: Actions (Week 3-4)
**Goal**: Port all poker actions with verification and execution logic

- [ ] Implement BaseAction interface
- [ ] Implement player actions:
  - [ ] FoldAction
  - [ ] CallAction
  - [ ] BetAction
  - [ ] RaiseAction
  - [ ] CheckAction
  - [ ] AllInAction
  - [ ] ShowAction
  - [ ] MuckAction
- [ ] Implement non-player actions:
  - [ ] DealAction
  - [ ] NewHandAction
  - [ ] JoinAction
  - [ ] LeaveAction
  - [ ] SitInAction
  - [ ] SitOutAction
- [ ] Write unit tests for each action

### Phase 4: Texas Hold'em (Week 4-5)
**Goal**: Implement complete Texas Hold'em game

- [ ] Implement TexasHoldem game struct
- [ ] Implement deal() method
- [ ] Implement reInit() method
- [ ] Implement performAction() method
- [ ] Implement getLegalActions() method
- [ ] Implement getNextPlayerToAct() method
- [ ] Implement round progression logic
- [ ] Implement winner calculation (integrate with hand evaluator)
- [ ] Write integration tests

### Phase 5: Hand Evaluation (Week 5)
**Goal**: Port or integrate poker hand evaluation

**Options**:
1. Port existing hand evaluator from TypeScript
2. Use existing Go poker hand evaluator library (e.g., `github.com/loganjspears/joker`)
3. Integrate with `pokersolver` if available in Go

**Decision**: Use `github.com/loganjspears/joker` for hand evaluation to save time

- [ ] Integrate hand evaluator library
- [ ] Adapt to internal Card representation
- [ ] Write tests for hand evaluation
- [ ] Implement winner calculation logic

### Phase 6: RPC Layer (Week 6)
**Goal**: Implement HTTP/JSON-RPC server

- [ ] Implement RPC request/response types
- [ ] Implement RPC handler
- [ ] Implement game state serialization
- [ ] Add CORS support
- [ ] Add health check endpoint
- [ ] Write RPC integration tests

### Phase 7: Testing & Validation (Week 6-7)
**Goal**: Ensure correctness and compatibility

- [ ] Port all TypeScript test cases
- [ ] Run full integration test suite
- [ ] Test edge cases (heads-up, sit-and-go, rake, etc.)
- [ ] Performance testing and optimization
- [ ] Memory profiling
- [ ] Compare results with TypeScript implementation

### Phase 8: Advanced Features (Week 7-8)
**Goal**: Add remaining features and optimizations

- [ ] Implement sit-and-go tournament logic
- [ ] Implement rake calculation
- [ ] Add ante support
- [ ] Implement auto-fold on timeout
- [ ] Add concurrent game support
- [ ] Performance optimizations
- [ ] Documentation

## Key Considerations

### Go-Specific Adaptations

1. **BigInt for Chip Amounts**
   - TypeScript uses `bigint` primitive
   - Go will use `*big.Int` from `math/big`
   - Be careful with pointer semantics and copying

2. **Error Handling**
   - TypeScript throws exceptions
   - Go returns errors explicitly
   - All methods that can fail should return `error`

3. **Interfaces vs Classes**
   - TypeScript uses classes with inheritance
   - Go uses interfaces and composition
   - Prefer embedding over inheritance

4. **Concurrency**
   - Add mutex locks for concurrent access
   - Consider using channels for event handling
   - Make game state thread-safe

5. **Memory Management**
   - Use pointers carefully
   - Consider memory pooling for frequently allocated objects
   - Profile memory usage

### Testing Strategy

1. **Unit Tests**
   - Test each action in isolation
   - Test each manager independently
   - Use table-driven tests where possible

2. **Integration Tests**
   - Port existing TypeScript test scenarios
   - Test full game flows
   - Test edge cases

3. **Compatibility Tests**
   - Run same test scenarios in both TS and Go
   - Compare game state at each step
   - Ensure deterministic behavior

### Performance Goals

- **Throughput**: 10,000+ actions per second
- **Latency**: < 1ms per action
- **Memory**: < 10MB per game instance
- **Concurrency**: Support 1000+ concurrent games

## Dependencies

### Required Go Packages

```go
// Standard library
- math/big          // For chip amounts
- encoding/json     // For JSON serialization
- net/http          // For RPC server
- crypto/sha256     // For cryptographic operations
- testing           // For tests

// Third-party
- github.com/loganjspears/joker  // Poker hand evaluation (optional)
- github.com/stretchr/testify    // Enhanced testing utilities
```

## Migration from TypeScript

### Type Mappings

| TypeScript | Go |
|------------|-----|
| `bigint` | `*big.Int` |
| `string` | `string` |
| `number` | `int` or `int64` |
| `boolean` | `bool` |
| `Map<K,V>` | `map[K]V` |
| `Array<T>` | `[]T` |
| `undefined` | Use zero values or pointers |
| `class` | `struct` with methods |
| `interface` | `interface` |
| `enum` | `const` with custom type |

### Common Patterns

**TypeScript Class**:
```typescript
class Player {
    constructor(
        public address: string,
        public chips: bigint
    ) {}

    bet(amount: bigint): void {
        this.chips -= amount;
    }
}
```

**Go Equivalent**:
```go
type Player struct {
    Address string
    Chips   *big.Int
}

func NewPlayer(address string, chips *big.Int) *Player {
    return &Player{
        Address: address,
        Chips:   chips,
    }
}

func (p *Player) Bet(amount *big.Int) {
    p.Chips = new(big.Int).Sub(p.Chips, amount)
}
```

## Success Criteria

1. **Functional Equivalence**: All TypeScript test cases pass in Go
2. **Performance**: Meet or exceed performance goals
3. **Code Quality**: Pass Go linting and formatting standards
4. **Documentation**: Comprehensive godoc comments
5. **Testing**: > 90% code coverage

## Timeline

- **Phase 1**: Week 1 (Foundation)
- **Phase 2**: Week 2-3 (Core Engine)
- **Phase 3**: Week 3-4 (Actions)
- **Phase 4**: Week 4-5 (Texas Hold'em)
- **Phase 5**: Week 5 (Hand Evaluation)
- **Phase 6**: Week 6 (RPC Layer)
- **Phase 7**: Week 6-7 (Testing & Validation)
- **Phase 8**: Week 7-8 (Advanced Features)

**Total**: 7-8 weeks for complete port

## Next Steps

1. Review and approve this approach
2. Start Phase 1 implementation
3. Set up CI/CD pipeline
4. Establish code review process
5. Begin porting utility functions and tests
