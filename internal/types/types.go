package types

import "math/big"

// PlayerActionType represents actions that players can take
type PlayerActionType string

const (
	ActionFold     PlayerActionType = "FOLD"
	ActionCheck    PlayerActionType = "CHECK"
	ActionCall     PlayerActionType = "CALL"
	ActionBet      PlayerActionType = "BET"
	ActionRaise    PlayerActionType = "RAISE"
	ActionAllIn    PlayerActionType = "ALL_IN"
	ActionMuck     PlayerActionType = "MUCK"
	ActionShow     PlayerActionType = "SHOW"
)

// NonPlayerActionType represents system actions
type NonPlayerActionType string

const (
	ActionDeal     NonPlayerActionType = "DEAL"
	ActionNewHand  NonPlayerActionType = "NEW_HAND"
	ActionJoin     NonPlayerActionType = "JOIN"
	ActionLeave    NonPlayerActionType = "LEAVE"
	ActionSitIn    NonPlayerActionType = "SIT_IN"
	ActionSitOut   NonPlayerActionType = "SIT_OUT"
)

// PlayerStatus represents the status of a player in the game
type PlayerStatus string

const (
	StatusActive     PlayerStatus = "ACTIVE"
	StatusFolded     PlayerStatus = "FOLDED"
	StatusAllIn      PlayerStatus = "ALL_IN"
	StatusSittingOut PlayerStatus = "SITTING_OUT"
	StatusBusted     PlayerStatus = "BUSTED"
)

// TexasHoldemRound represents the current round of betting
type TexasHoldemRound string

const (
	RoundPreFlop  TexasHoldemRound = "PRE_FLOP"
	RoundFlop     TexasHoldemRound = "FLOP"
	RoundTurn     TexasHoldemRound = "TURN"
	RoundRiver    TexasHoldemRound = "RIVER"
	RoundShowdown TexasHoldemRound = "SHOWDOWN"
)

// GameFormat represents the format of the poker game
type GameFormat string

const (
	FormatCash       GameFormat = "CASH"
	FormatSitAndGo   GameFormat = "SIT_AND_GO"
	FormatTournament GameFormat = "TOURNAMENT"
)

// GameVariant represents the poker variant being played
type GameVariant string

const (
	VariantTexasHoldem GameVariant = "TEXAS_HOLDEM"
	VariantOmaha       GameVariant = "OMAHA"
)

// Suit represents a card suit (matches TypeScript SDK SUIT enum)
type Suit int

const (
	SuitClubs    Suit = 1
	SuitDiamonds Suit = 2
	SuitHearts   Suit = 3
	SuitSpades   Suit = 4
)

// Card represents a playing card (matches TypeScript SDK Card type)
type Card struct {
	Suit     Suit   // 1=Clubs, 2=Diamonds, 3=Hearts, 4=Spades
	Rank     int    // 1-13 where 1=Ace, 10=Ten, 11=Jack, 12=Queen, 13=King
	Value    int    // Calculated as 13 * (suit - 1) + (rank - 1)
	Mnemonic string // e.g., "AS", "2C", "KH", "TD"
}

// Range represents the min and max betting amounts
type Range struct {
	MinAmount *big.Int
	MaxAmount *big.Int
}

// Turn represents a player's action in the game
type Turn struct {
	PlayerID string
	Action   interface{} // PlayerActionType or NonPlayerActionType
	Amount   *big.Int
	Index    int
}

// TurnWithSeat extends Turn with seat and timestamp information
type TurnWithSeat struct {
	Turn
	Seat      int
	Timestamp int64
}

// Winner represents a winner of a hand
type Winner struct {
	Amount      *big.Int
	Cards       []string
	Name        string
	Description string
}

// GameOptions represents the configuration options for a poker game
type GameOptions struct {
	Format         GameFormat
	Variant        GameVariant
	SmallBlind     *big.Int
	BigBlind       *big.Int
	MinPlayers     int
	MaxPlayers     int
	Ante           *big.Int
	RakePercentage float64
}
