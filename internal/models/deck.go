package models

import (
	"errors"
	"math/rand"
	"time"

	"github.com/block52/go-pvm/internal/types"
)

// Deck represents a deck of playing cards
type Deck struct {
	cards    []types.Card
	position int
}

var (
	ranks = []string{"2", "3", "4", "5", "6", "7", "8", "9", "T", "J", "Q", "K", "A"}
	suits = []string{"c", "d", "h", "s"} // clubs, diamonds, hearts, spades
)

// NewDeck creates a new standard 52-card deck
func NewDeck() *Deck {
	cards := make([]types.Card, 0, 52)
	for _, suit := range suits {
		for _, rank := range ranks {
			cards = append(cards, types.Card{Rank: rank, Suit: suit})
		}
	}

	return &Deck{
		cards:    cards,
		position: 0,
	}
}

// Shuffle randomizes the order of cards in the deck
func (d *Deck) Shuffle() {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	r.Shuffle(len(d.cards), func(i, j int) {
		d.cards[i], d.cards[j] = d.cards[j], d.cards[i]
	})
	d.position = 0
}

// GetNext returns the next card from the deck
func (d *Deck) GetNext() (types.Card, error) {
	if d.position >= len(d.cards) {
		return types.Card{}, errors.New("no more cards in deck")
	}
	card := d.cards[d.position]
	d.position++
	return card, nil
}

// Reset resets the deck position to the beginning
func (d *Deck) Reset() {
	d.position = 0
}

// Remaining returns the number of cards remaining in the deck
func (d *Deck) Remaining() int {
	return len(d.cards) - d.position
}
