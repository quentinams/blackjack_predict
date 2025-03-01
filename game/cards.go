package game

import "math/rand"

type Rank string
type Suit string

const (
	Spades   Suit = "Spades"
	Hearts   Suit = "Hearts"
	Diamonds Suit = "Diamonds"
	Clubs    Suit = "Clubs"
)

const (
	Ace   Rank = "As"
	Two   Rank = "2"
	Three Rank = "3"
	Four  Rank = "4"
	Five  Rank = "5"
	Six   Rank = "6"
	Seven Rank = "7"
	Eight Rank = "8"
	Nine  Rank = "9"
	Ten   Rank = "10"
	Jack  Rank = "Jack"
	Queen Rank = "Queen"
	King  Rank = "King"
)

type Card struct {
	Suit  Suit
	Rank  Rank
	Value int
}

// function qui mélange les cartes
func Shuffle(deck []Card) []Card {
	for i := range deck {
		j := rand.Intn(i + 1)
		deck[i], deck[j] = deck[j], deck[i]
	}
	return deck
}

func NewDeck() []Card {
	suits := []Suit{Spades, Hearts, Diamonds, Clubs}
	ranks := []Rank{Ace, Two, Three, Four, Five, Six, Seven, Eight, Nine, Ten, Jack, Queen, King}
	value := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 10, 10, 10}
	deck := make([]Card, 0, 52)

	for _, suit := range suits {
		for _, rank := range ranks {
			deck = append(deck, Card{Suit: suit, Rank: rank, Value: value[len(deck)%13]})

		}
	}

	return Shuffle(deck)
}

func Draw(deck *[]Card) Card {
	card := (*deck)[0]
	*deck = (*deck)[1:]
	return card
}
