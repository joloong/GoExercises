package deck

import (
	"fmt"
	"testing"
)

func ExampleCard() {
	fmt.Println(Card{Rank: Ace, Suit: Heart})
	fmt.Println(Card{Rank: Two, Suit: Spade})
	fmt.Println(Card{Rank: Ten, Suit: Diamond})
	fmt.Println(Card{Rank: King, Suit: Club})
	fmt.Println(Card{Suit: Joker})

	// Output:
	// Ace of Hearts
	// Two of Spades
	// Ten of Diamonds
	// King of Clubs
	// Joker
}

func TestNew(t *testing.T) {
	cards := New()
	// 4 Suits of 13 Ranks = 52 Cards
	if len(cards) != 52 {
		t.Error("Wrong number of cards in a new deck.")
	}
}

func TestDefaultSort(t *testing.T) {
	cards := New(DefaultSort)
	want := Card{Rank: Ace, Suit: Spade}
	if cards[0] != want {
		t.Error("Expected Ace of Spades as first card. Received: ", cards[0])
	}
	dWant := Card{Rank: King, Suit: Heart}
	if cards[len(cards)-1] != dWant {
		t.Error("Expected King of Hearts as last card. Received: ", cards[len(cards)-1])
	}
}

func TestSort(t *testing.T) {
	cards := New(Sort(Less))
	exp := Card{Rank: Ace, Suit: Spade}
	if cards[0] != exp {
		t.Error("Expected Ace of Spades as first card. Received: ", cards[0])
	}
}

func TestJokers(t *testing.T) {
	numOfJoker := 3
	cards := New(Jokers(numOfJoker))
	count := 0
	for _, c := range cards {
		if c.Suit == Joker {
			count++
		}
	}

	if count != numOfJoker {
		t.Errorf("Expected %d Jokers. Received: %d", numOfJoker, count)
	}
}

func TestFilter(t *testing.T) {
	filter := func(card Card) bool {
		return card.Rank == Two || card.Rank == Three
	}

	cards := New(Filter(filter))
	for _, c := range cards {
		if c.Rank == Two || c.Rank == Three {
			t.Error("Expected all Two(s) and Three(s) to be filtered out.")
		}
	}
}

func TestDeck(t *testing.T) {
	cards := New(Deck(3))
	if len(cards) != 13*4*3 {
		t.Errorf("Expected %d cards, received %d cards.", 13*4*3, len(cards))
	}
}
