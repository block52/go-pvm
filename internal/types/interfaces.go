package types

import "math/big"

// IAction defines the interface for poker actions
type IAction interface {
	Type() interface{} // Returns PlayerActionType or NonPlayerActionType
	Verify(player IPlayer) (*Range, error)
	Execute(player IPlayer, index int, amount *big.Int) error
}

// IPlayer defines the interface for a player
type IPlayer interface {
	GetAddress() string
	GetChips() *big.Int
	GetStatus() PlayerStatus
	GetCards() []Card
	SetChips(chips *big.Int)
	SetStatus(status PlayerStatus)
	SetCards(cards []Card)
}

// IPoker defines the core poker game interface
type IPoker interface {
	// Game configuration
	GetGameFormat() GameFormat
	GetGameVariant() GameVariant
	GetSmallBlind() *big.Int
	GetBigBlind() *big.Int

	// Game flow
	Deal()
	ReInit(deck string) error
	HasRoundEnded(round TexasHoldemRound) bool

	// Player queries
	GetLegalActions(address string) ([]LegalActionDTO, error)
	GetNextPlayerToAct() (IPlayer, error)
	GetPlayersLastAction(address string) (*TurnWithSeat, error)

	// Action management
	GetActionIndex() int
	GetLastRoundAction() (*Turn, error)
	PerformAction(address string, action PlayerActionType, index int, amount *big.Int) error

	// Betting/Pot
	GetBets(round TexasHoldemRound) map[string]*big.Int
	GetPot() *big.Int
}

// IDealer defines what the dealer position manager needs
type IDealer interface {
	GetLastActedSeat() int
	GetDealerPosition() int
	GetMinPlayers() int
	GetMaxPlayers() int
	FindActivePlayers() []IPlayer
	GetPlayerAtSeat(seat int) (IPlayer, error)
	GetPlayerSeatNumber(playerID string) int
}

// IDealerPositionManager defines the dealer position management interface
type IDealerPositionManager interface {
	GetDealerPosition() int
	HandlePlayerLeave(seat int)
	HandlePlayerJoin(seat int)
	HandleNewHand() int
	GetPosition(name string) int
	GetSmallBlindPosition() int
	GetBigBlindPosition() int
	ValidateDealerPosition() bool
	FindNextActivePlayer(currentSeat int) (IPlayer, error)
}

// LegalActionDTO represents a legal action a player can take
type LegalActionDTO struct {
	Action   PlayerActionType
	MinAmount *big.Int
	MaxAmount *big.Int
}
