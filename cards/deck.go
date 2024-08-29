package cards

import "fmt"

// NewDeck crée un nouveau jeu de 52 cartes.
func NewDeck() []Card {
	ranks := []Card{As, Deux, Trois, Quatre, Cinq, Six, Sept, Huit, Neuf, Dix, Valet, Dame, Roi}

	deck := make([]Card, 0, 52)

	for _, rank := range ranks {
		deck = append(deck, rank)
	}

	return deck
}

func PrintDeck(cards []Card) {
	for _, card := range cards {
		fmt.Println(card)
	}
}
