package models

import (
	"testing"

	"github.com/block52/go-pvm/internal/types"
)

func TestNewDeck_Standard52(t *testing.T) {
	deck, err := NewDeck("")
	if err != nil {
		t.Fatalf("NewDeck failed: %v", err)
	}

	// Should have 52 cards
	if len(deck.cards) != 52 {
		t.Errorf("Expected 52 cards, got %d", len(deck.cards))
	}

	// Should start at position 0
	if deck.top != 0 {
		t.Errorf("Expected top=0, got %d", deck.top)
	}

	// Should have a hash
	if deck.hash == "" {
		t.Error("Expected deck to have a hash")
	}

	// Verify first card is Ace of Clubs (rank=1, suit=1)
	if deck.cards[0].Rank != 1 || deck.cards[0].Suit != types.SuitClubs {
		t.Errorf("Expected first card to be Ace of Clubs, got rank=%d suit=%d", deck.cards[0].Rank, deck.cards[0].Suit)
	}
	if deck.cards[0].Mnemonic != "AC" {
		t.Errorf("Expected first card mnemonic to be 'AC', got '%s'", deck.cards[0].Mnemonic)
	}
	if deck.cards[0].Value != 0 {
		t.Errorf("Expected first card value to be 0, got %d", deck.cards[0].Value)
	}

	// Verify last card is King of Spades (rank=13, suit=4)
	lastCard := deck.cards[51]
	if lastCard.Rank != 13 || lastCard.Suit != types.SuitSpades {
		t.Errorf("Expected last card to be King of Spades, got rank=%d suit=%d", lastCard.Rank, lastCard.Suit)
	}
	if lastCard.Mnemonic != "KS" {
		t.Errorf("Expected last card mnemonic to be 'KS', got '%s'", lastCard.Mnemonic)
	}
	// Value = 13 * (4 - 1) + (13 - 1) = 39 + 12 = 51
	if lastCard.Value != 51 {
		t.Errorf("Expected last card value to be 51, got %d", lastCard.Value)
	}
}

func TestGetCardMnemonic(t *testing.T) {
	tests := []struct {
		suit     types.Suit
		rank     int
		expected string
	}{
		{types.SuitClubs, 1, "AC"},
		{types.SuitDiamonds, 2, "2D"},
		{types.SuitHearts, 10, "TH"},
		{types.SuitSpades, 11, "JS"},
		{types.SuitClubs, 12, "QC"},
		{types.SuitDiamonds, 13, "KD"},
		{types.SuitHearts, 9, "9H"},
	}

	for _, tt := range tests {
		result := GetCardMnemonic(tt.suit, tt.rank)
		if result != tt.expected {
			t.Errorf("GetCardMnemonic(%d, %d) = %s; want %s", tt.suit, tt.rank, result, tt.expected)
		}
	}
}

func TestFromString(t *testing.T) {
	tests := []struct {
		mnemonic     string
		expectedRank int
		expectedSuit types.Suit
		shouldError  bool
	}{
		{"AS", 1, types.SuitSpades, false},
		{"2C", 2, types.SuitClubs, false},
		{"10H", 10, types.SuitHearts, false},
		{"TD", 10, types.SuitDiamonds, false},
		{"JH", 11, types.SuitHearts, false},
		{"QS", 12, types.SuitSpades, false},
		{"KC", 13, types.SuitClubs, false},
		{"as", 1, types.SuitSpades, false}, // lowercase should work
		{"invalid", 0, 0, true},
		{"1X", 0, 0, true},
	}

	for _, tt := range tests {
		card, err := FromString(tt.mnemonic)
		if tt.shouldError {
			if err == nil {
				t.Errorf("FromString(%s) expected error, got nil", tt.mnemonic)
			}
			continue
		}

		if err != nil {
			t.Errorf("FromString(%s) unexpected error: %v", tt.mnemonic, err)
			continue
		}

		if card.Rank != tt.expectedRank {
			t.Errorf("FromString(%s) rank = %d; want %d", tt.mnemonic, card.Rank, tt.expectedRank)
		}
		if card.Suit != tt.expectedSuit {
			t.Errorf("FromString(%s) suit = %d; want %d", tt.mnemonic, card.Suit, tt.expectedSuit)
		}

		// Verify value calculation
		expectedValue := 13*(int(tt.expectedSuit)-1) + (tt.expectedRank - 1)
		if card.Value != expectedValue {
			t.Errorf("FromString(%s) value = %d; want %d", tt.mnemonic, card.Value, expectedValue)
		}
	}
}

func TestGetNext(t *testing.T) {
	deck, _ := NewDeck("")

	// Get first card
	card, err := deck.GetNext()
	if err != nil {
		t.Fatalf("GetNext failed: %v", err)
	}

	// Should be Ace of Clubs
	if card.Mnemonic != "AC" {
		t.Errorf("Expected first card to be AC, got %s", card.Mnemonic)
	}

	// Position should be 1
	if deck.top != 1 {
		t.Errorf("Expected top=1, got %d", deck.top)
	}

	// Deal all remaining cards
	for i := 1; i < 52; i++ {
		_, err := deck.GetNext()
		if err != nil {
			t.Fatalf("GetNext failed at position %d: %v", i, err)
		}
	}

	// Should error when no more cards
	_, err = deck.GetNext()
	if err == nil {
		t.Error("Expected error when getting card from empty deck")
	}
}

func TestDeal(t *testing.T) {
	deck, _ := NewDeck("")

	// Deal 5 cards
	cards, err := deck.Deal(5)
	if err != nil {
		t.Fatalf("Deal failed: %v", err)
	}

	if len(cards) != 5 {
		t.Errorf("Expected 5 cards, got %d", len(cards))
	}

	// Position should be 5
	if deck.top != 5 {
		t.Errorf("Expected top=5, got %d", deck.top)
	}

	// Deal remaining 47 cards
	_, err = deck.Deal(47)
	if err != nil {
		t.Errorf("Expected to deal 47 cards successfully, got error: %v", err)
	}

	// Try to deal when empty
	_, err = deck.Deal(1)
	if err == nil {
		t.Error("Expected error when dealing from empty deck")
	}
}

func TestToString(t *testing.T) {
	deck, _ := NewDeck("")

	// Initial string should have [AC] at position 0
	str := deck.ToString()
	if str[:4] != "[AC]" {
		t.Errorf("Expected string to start with '[AC]', got '%s'", str[:4])
	}

	// Deal 2 cards
	deck.GetNext()
	deck.GetNext()

	// Now position 2 should be marked
	str = deck.ToString()
	if str[:8] != "AC-2C-[3" {
		t.Errorf("Expected string to start with 'AC-2C-[3', got '%s'", str[:8])
	}
}

func TestNewDeck_FromString(t *testing.T) {
	// Create a standard deck
	deck1, _ := NewDeck("")

	// Get its string representation
	deckStr := deck1.ToString()

	// Create a new deck from the string
	deck2, err := NewDeck(deckStr)
	if err != nil {
		t.Fatalf("NewDeck from string failed: %v", err)
	}

	// Both decks should be identical
	if len(deck1.cards) != len(deck2.cards) {
		t.Errorf("Card count mismatch: %d vs %d", len(deck1.cards), len(deck2.cards))
	}

	// Top position should match
	if deck1.top != deck2.top {
		t.Errorf("Top position mismatch: %d vs %d", deck1.top, deck2.top)
	}

	// Hashes should match
	if deck1.hash != deck2.hash {
		t.Errorf("Hash mismatch: %s vs %s", deck1.hash, deck2.hash)
	}
}

func TestNewDeck_FromStringWithPosition(t *testing.T) {
	// Deck string with position marker at card 10
	deckStr := "AC-2C-3C-4C-5C-6C-7C-8C-9C-[TC]-JC-QC-KC-AD-2D-3D-4D-5D-6D-7D-8D-9D-TD-JD-QD-KD-AH-2H-3H-4H-5H-6H-7H-8H-9H-TH-JH-QH-KH-AS-2S-3S-4S-5S-6S-7S-8S-9S-TS-JS-QS-KS"

	deck, err := NewDeck(deckStr)
	if err != nil {
		t.Fatalf("NewDeck from string failed: %v", err)
	}

	// Position should be 9 (0-indexed)
	if deck.top != 9 {
		t.Errorf("Expected top=9, got %d", deck.top)
	}

	// Next card should be TC
	card, _ := deck.GetNext()
	if card.Mnemonic != "TC" {
		t.Errorf("Expected next card to be TC, got %s", card.Mnemonic)
	}
}

func TestCardValueCalculation(t *testing.T) {
	tests := []struct {
		suit          types.Suit
		rank          int
		expectedValue int
	}{
		{types.SuitClubs, 1, 0},      // AC: 13*(1-1) + (1-1) = 0
		{types.SuitClubs, 13, 12},    // KC: 13*(1-1) + (13-1) = 12
		{types.SuitDiamonds, 1, 13},  // AD: 13*(2-1) + (1-1) = 13
		{types.SuitHearts, 1, 26},    // AH: 13*(3-1) + (1-1) = 26
		{types.SuitSpades, 1, 39},    // AS: 13*(4-1) + (1-1) = 39
		{types.SuitSpades, 13, 51},   // KS: 13*(4-1) + (13-1) = 51
	}

	for _, tt := range tests {
		mnemonic := GetCardMnemonic(tt.suit, tt.rank)
		card, err := FromString(mnemonic)
		if err != nil {
			t.Fatalf("FromString(%s) failed: %v", mnemonic, err)
		}

		if card.Value != tt.expectedValue {
			t.Errorf("Card %s: expected value %d, got %d", mnemonic, tt.expectedValue, card.Value)
		}
	}
}

func TestRemaining(t *testing.T) {
	deck, _ := NewDeck("")

	if deck.Remaining() != 52 {
		t.Errorf("Expected 52 remaining, got %d", deck.Remaining())
	}

	deck.Deal(10)
	if deck.Remaining() != 42 {
		t.Errorf("Expected 42 remaining, got %d", deck.Remaining())
	}

	deck.Deal(42)
	if deck.Remaining() != 0 {
		t.Errorf("Expected 0 remaining, got %d", deck.Remaining())
	}
}
