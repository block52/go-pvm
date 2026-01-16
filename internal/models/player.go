package models

import (
	"math/big"

	"github.com/block52/go-pvm/internal/types"
)

// Player represents a poker player
type Player struct {
	Address   string
	Chips     *big.Int
	Status    types.PlayerStatus
	HoleCards []types.Card
	Seat      int
}

// NewPlayer creates a new player instance
func NewPlayer(address string, chips *big.Int, seat int) *Player {
	return &Player{
		Address:   address,
		Chips:     chips,
		Status:    types.StatusActive,
		HoleCards: make([]types.Card, 0),
		Seat:      seat,
	}
}

// GetAddress returns the player's address
func (p *Player) GetAddress() string {
	return p.Address
}

// GetChips returns the player's chip count
func (p *Player) GetChips() *big.Int {
	return p.Chips
}

// GetStatus returns the player's status
func (p *Player) GetStatus() types.PlayerStatus {
	return p.Status
}

// GetCards returns the player's hole cards
func (p *Player) GetCards() []types.Card {
	return p.HoleCards
}

// SetChips sets the player's chip count
func (p *Player) SetChips(chips *big.Int) {
	p.Chips = chips
}

// SetStatus sets the player's status
func (p *Player) SetStatus(status types.PlayerStatus) {
	p.Status = status
}

// SetCards sets the player's hole cards
func (p *Player) SetCards(cards []types.Card) {
	p.HoleCards = cards
}
