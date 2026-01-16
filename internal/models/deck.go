package models

import (
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"github.com/block52/go-pvm/internal/types"
)

// Deck represents a deck of playing cards
// Matches TypeScript Deck implementation from pvm/ts/src/models/deck.ts
type Deck struct {
	cards []types.Card
	hash  string
	top   int // Current position in deck (like TypeScript 'top')
}

// NewDeck creates a new deck from an optional deck string
// If deckStr is empty, creates a standard 52-card deck
// Matches TypeScript constructor
func NewDeck(deckStr string) (*Deck, error) {
	d := &Deck{
		cards: make([]types.Card, 0, 52),
		hash:  "",
		top:   0,
	}

	// For backwards compatibility: treat empty strings as undefined (create standard deck)
	deckStr = strings.TrimSpace(deckStr)

	if deckStr != "" {
		// Parse deck from string format like "AS-2C-3D-[4H]-5S-..."
		mnemonics := strings.Split(deckStr, "-")
		if len(mnemonics) != 52 {
			return nil, errors.New("deck must contain 52 cards")
		}

		for i, mnemonic := range mnemonics {
			// Check if this is the current top position (marked with brackets)
			if strings.HasPrefix(mnemonic, "[") && strings.HasSuffix(mnemonic, "]") {
				mnemonic = strings.Trim(mnemonic, "[]")
				d.top = i
			}

			card, err := FromString(mnemonic)
			if err != nil {
				return nil, fmt.Errorf("invalid card at position %d: %w", i, err)
			}
			d.cards = append(d.cards, card)
		}
	} else {
		d.initStandard52()
	}

	d.createHash()
	return d, nil
}

// GetNext returns the next card from the deck
// Matches TypeScript getNext()
func (d *Deck) GetNext() (types.Card, error) {
	if d.top >= len(d.cards) {
		return types.Card{}, errors.New("no more cards in deck")
	}
	card := d.cards[d.top]
	d.top++
	return card, nil
}

// Deal deals a specified number of cards
// Matches TypeScript deal(amount)
func (d *Deck) Deal(amount int) ([]types.Card, error) {
	if d.top+amount > len(d.cards) {
		return nil, errors.New("not enough cards in deck")
	}

	cards := make([]types.Card, amount)
	for i := 0; i < amount; i++ {
		card, err := d.GetNext()
		if err != nil {
			return nil, err
		}
		cards[i] = card
	}
	return cards, nil
}

// ToString serializes the deck to string format with position marker
// Matches TypeScript toString()
func (d *Deck) ToString() string {
	mnemonics := make([]string, len(d.cards))
	for i, card := range d.cards {
		if i == d.top {
			mnemonics[i] = fmt.Sprintf("[%s]", card.Mnemonic)
		} else {
			mnemonics[i] = card.Mnemonic
		}
	}
	return strings.Join(mnemonics, "-")
}

// GetHash returns the SHA256 hash of the deck
func (d *Deck) GetHash() string {
	return d.hash
}

// GetTop returns the current position in the deck
func (d *Deck) GetTop() int {
	return d.top
}

// Remaining returns the number of cards remaining in the deck
func (d *Deck) Remaining() int {
	return len(d.cards) - d.top
}

// createHash generates SHA256 hash of the deck
// Matches TypeScript createHash()
func (d *Deck) createHash() {
	mnemonics := make([]string, len(d.cards))
	for i, card := range d.cards {
		mnemonics[i] = card.Mnemonic
	}
	cardsAsString := strings.Join(mnemonics, "-")
	hash := sha256.Sum256([]byte(cardsAsString))
	d.hash = hex.EncodeToString(hash[:])
}

// initStandard52 initializes a standard 52-card deck
// Matches TypeScript initStandard52()
func (d *Deck) initStandard52() {
	d.cards = make([]types.Card, 0, 52)

	// Iterate through suits: CLUBS=1, DIAMONDS=2, HEARTS=3, SPADES=4
	for suit := types.SuitClubs; suit <= types.SuitSpades; suit++ {
		// Iterate through ranks: 1=Ace, 2-9=numbers, 10=Ten, 11=Jack, 12=Queen, 13=King
		for rank := 1; rank <= 13; rank++ {
			mnemonic := GetCardMnemonic(suit, rank)
			value := 13*(int(suit)-1) + (rank - 1)
			d.cards = append(d.cards, types.Card{
				Suit:     suit,
				Rank:     rank,
				Value:    value,
				Mnemonic: mnemonic,
			})
		}
	}
}

// GetCardMnemonic generates the mnemonic string for a card
// Matches TypeScript getCardMnemonic()
func GetCardMnemonic(suit types.Suit, rank int) string {
	// Map special ranks
	rankStr := ""
	switch rank {
	case 1:
		rankStr = "A"
	case 10:
		rankStr = "T"
	case 11:
		rankStr = "J"
	case 12:
		rankStr = "Q"
	case 13:
		rankStr = "K"
	default:
		rankStr = strconv.Itoa(rank)
	}

	// Map suit to string
	suitStr := ""
	switch suit {
	case types.SuitClubs:
		suitStr = "C"
	case types.SuitDiamonds:
		suitStr = "D"
	case types.SuitHearts:
		suitStr = "H"
	case types.SuitSpades:
		suitStr = "S"
	}

	return rankStr + suitStr
}

// FromString parses a card mnemonic string into a Card
// Matches TypeScript fromString()
func FromString(mnemonic string) (types.Card, error) {
	// Match pattern like "AS", "2C", "10H", "KD"
	re := regexp.MustCompile(`^([AJQKTajqkt]|[0-9]+)([CDHS])$`)
	matches := re.FindStringSubmatch(strings.ToUpper(mnemonic))

	if matches == nil {
		return types.Card{}, fmt.Errorf("invalid card mnemonic: %s", mnemonic)
	}

	rankStr := matches[1]
	suitChar := matches[2]

	// Convert rank string to number
	var rank int
	switch rankStr {
	case "A":
		rank = 1
	case "T":
		rank = 10
	case "J":
		rank = 11
	case "Q":
		rank = 12
	case "K":
		rank = 13
	default:
		var err error
		rank, err = strconv.Atoi(rankStr)
		if err != nil {
			return types.Card{}, fmt.Errorf("invalid rank: %s", rankStr)
		}
	}

	// Convert suit character to Suit enum
	var suit types.Suit
	switch suitChar {
	case "C":
		suit = types.SuitClubs
	case "D":
		suit = types.SuitDiamonds
	case "H":
		suit = types.SuitHearts
	case "S":
		suit = types.SuitSpades
	default:
		return types.Card{}, fmt.Errorf("invalid suit character: %s", suitChar)
	}

	// Calculate value: 13 * (suit - 1) + (rank - 1)
	value := 13*(int(suit)-1) + (rank - 1)

	return types.Card{
		Suit:     suit,
		Rank:     rank,
		Value:    value,
		Mnemonic: strings.ToUpper(mnemonic), // Use original mnemonic for consistency
	}, nil
}
