package models

import (
	"testing"

	"github.com/block52/go-pvm/internal/types"
)

// TestDeck_Constructor tests the Deck constructor behavior
func TestDeck_Constructor(t *testing.T) {
	t.Run("should initialize with default values", func(t *testing.T) {
		deck, err := NewDeck("")
		if err != nil {
			t.Fatalf("NewDeck failed: %v", err)
		}

		if deck.hash == "" {
			t.Error("Expected deck.hash to be defined")
		}
	})

	t.Run("should initialize with standard 52-card deck", func(t *testing.T) {
		mnemonic := "AC-2C-3C-4C-5C-6C-7C-8C-9C-TC-JC-QC-KC-" +
			"AD-2D-3D-4D-5D-6D-7D-8D-9D-TD-JD-QD-KD-" +
			"AH-2H-3H-4H-5H-6H-7H-8H-9H-TH-JH-QH-KH-" +
			"AS-2S-3S-4S-5S-6S-7S-8S-9S-TS-JS-QS-KS"

		deck, err := NewDeck(mnemonic)
		if err != nil {
			t.Fatalf("NewDeck failed: %v", err)
		}

		if len(deck.cards) != 52 {
			t.Errorf("Expected 52 cards, got %d", len(deck.cards))
		}
	})

	t.Run("should serialize to string", func(t *testing.T) {
		mnemonic := "[AC]-2C-3C-4C-5C-6C-7C-8C-9C-TC-JC-QC-KC-" +
			"AD-2D-3D-4D-5D-6D-7D-8D-9D-TD-JD-QD-KD-" +
			"AH-2H-3H-4H-5H-6H-7H-8H-9H-TH-JH-QH-KH-" +
			"AS-2S-3S-4S-5S-6S-7S-8S-9S-TS-JS-QS-KS"

		deck, err := NewDeck(mnemonic)
		if err != nil {
			t.Fatalf("NewDeck failed: %v", err)
		}

		result := deck.ToString()
		if result != mnemonic {
			t.Errorf("Expected toString to equal input mnemonic.\nGot:      %s\nExpected: %s", result, mnemonic)
		}
	})

	t.Run("should treat empty string as undefined (create standard deck)", func(t *testing.T) {
		emptyDeck, err := NewDeck("")
		if err != nil {
			t.Fatalf("NewDeck failed: %v", err)
		}

		if len(emptyDeck.cards) != 52 {
			t.Errorf("Expected 52 cards, got %d", len(emptyDeck.cards))
		}
	})

	t.Run("should treat whitespace-only string as undefined (create standard deck)", func(t *testing.T) {
		whitespaceDeck, err := NewDeck("   ")
		if err != nil {
			t.Fatalf("NewDeck failed: %v", err)
		}

		if len(whitespaceDeck.cards) != 52 {
			t.Errorf("Expected 52 cards, got %d", len(whitespaceDeck.cards))
		}
	})

	t.Run("should initialize standard deck if no parameter provided", func(t *testing.T) {
		standardDeck, err := NewDeck("")
		if err != nil {
			t.Fatalf("NewDeck failed: %v", err)
		}

		if len(standardDeck.cards) != 52 {
			t.Errorf("Expected 52 cards, got %d", len(standardDeck.cards))
		}
	})
}

// TestDeck_GetCardMnemonic tests the GetCardMnemonic function
func TestDeck_GetCardMnemonic(t *testing.T) {
	t.Run("should convert number cards correctly", func(t *testing.T) {
		result1 := GetCardMnemonic(types.SuitSpades, 2)
		if result1 != "2S" {
			t.Errorf("Expected '2S', got '%s'", result1)
		}

		result2 := GetCardMnemonic(types.SuitHearts, 10)
		if result2 != "TH" {
			t.Errorf("Expected 'TH', got '%s'", result2)
		}
	})

	t.Run("should convert face cards correctly", func(t *testing.T) {
		result1 := GetCardMnemonic(types.SuitClubs, 11)
		if result1 != "JC" {
			t.Errorf("Expected 'JC', got '%s'", result1)
		}

		result2 := GetCardMnemonic(types.SuitDiamonds, 12)
		if result2 != "QD" {
			t.Errorf("Expected 'QD', got '%s'", result2)
		}

		result3 := GetCardMnemonic(types.SuitHearts, 13)
		if result3 != "KH" {
			t.Errorf("Expected 'KH', got '%s'", result3)
		}

		result4 := GetCardMnemonic(types.SuitSpades, 1)
		if result4 != "AS" {
			t.Errorf("Expected 'AS', got '%s'", result4)
		}
	})
}

// TestDeck_GetNextAndDeal tests the GetNext and Deal methods
func TestDeck_GetNextAndDeal(t *testing.T) {
	t.Run("should draw next card correctly", func(t *testing.T) {
		deck, _ := NewDeck("")

		card, err := deck.GetNext()
		if err != nil {
			t.Fatalf("GetNext failed: %v", err)
		}

		if card.Suit == 0 {
			t.Error("Expected card.suit to be defined")
		}
		if card.Rank == 0 {
			t.Error("Expected card.rank to be defined")
		}
		if card.Mnemonic == "" {
			t.Error("Expected card.mnemonic to be defined")
		}

		nextCard, err := deck.GetNext()
		if err != nil {
			t.Fatalf("GetNext failed: %v", err)
		}

		if card.Mnemonic == nextCard.Mnemonic {
			t.Error("Expected card and nextCard to be different")
		}
	})

	t.Run("should deal multiple cards", func(t *testing.T) {
		deck, _ := NewDeck("")

		cards, err := deck.Deal(5)
		if err != nil {
			t.Fatalf("Deal failed: %v", err)
		}

		if len(cards) != 5 {
			t.Errorf("Expected 5 cards, got %d", len(cards))
		}

		for i, card := range cards {
			if card.Suit == 0 {
				t.Errorf("Card %d: expected suit to be defined", i)
			}
			if card.Rank == 0 {
				t.Errorf("Card %d: expected rank to be defined", i)
			}
			if card.Mnemonic == "" {
				t.Errorf("Card %d: expected mnemonic to be defined", i)
			}
		}

		// Check if top index has moved
		nextCard, err := deck.GetNext()
		if err != nil {
			t.Fatalf("GetNext failed: %v", err)
		}

		if cards[0].Mnemonic == nextCard.Mnemonic {
			t.Error("Expected nextCard to be different from first dealt card")
		}
	})
}

// TestDeck_ToJson tests the ToJson method
func TestDeck_ToJson(t *testing.T) {
	t.Run("should serialize deck state", func(t *testing.T) {
		deck, _ := NewDeck("")

		json := deck.ToJson()

		if json.Cards == nil {
			t.Error("Expected json to have 'Cards' property")
		}

		if len(json.Cards) == 0 {
			t.Error("Expected json.Cards to be a non-empty array")
		}
	})
}

// TestDeck_InitStandard52 tests the initStandard52 method
func TestDeck_InitStandard52(t *testing.T) {
	t.Run("should create a standard 52-card deck", func(t *testing.T) {
		deck, _ := NewDeck("")

		json := deck.ToJson()
		if len(json.Cards) != 52 {
			t.Errorf("Expected 52 cards, got %d", len(json.Cards))
		}

		// Check for Ace of Spades (rank 1)
		hasAceOfSpades := false
		for _, card := range json.Cards {
			if card.Suit == types.SuitSpades && card.Rank == 1 {
				hasAceOfSpades = true
				break
			}
		}

		if !hasAceOfSpades {
			t.Error("Expected deck to contain Ace of Spades")
		}
	})
}

// TestDeck_HashGeneration tests the hash generation
func TestDeck_HashGeneration(t *testing.T) {
	t.Run("should create different hashes for different card orders", func(t *testing.T) {
		standardDeck, _ := NewDeck("")
		standardHash := standardDeck.hash

		// Create a deck with a different order (reversed)
		reversedDeckStr := "KS-QS-JS-TS-9S-8S-7S-6S-5S-4S-3S-2S-AS-" +
			"KH-QH-JH-TH-9H-8H-7H-6H-5H-4H-3H-2H-AH-" +
			"KD-QD-JD-TD-9D-8D-7D-6D-5D-4D-3D-2D-AD-" +
			"KC-QC-JC-TC-9C-8C-7C-6C-5C-4C-3C-2C-AC"

		reversedDeck, err := NewDeck(reversedDeckStr)
		if err != nil {
			t.Fatalf("NewDeck failed: %v", err)
		}
		reversedHash := reversedDeck.hash

		// Different card orders should produce different hashes
		if reversedHash == standardHash {
			t.Error("Expected different hashes for different card orders")
		}
	})
}
