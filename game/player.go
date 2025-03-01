package game

import "fmt"

type Player struct {
	Cash int
	Hand []Card
}

func (p *Player) AddCard(card Card) {
	p.Hand = append(p.Hand, card)
}

func (p *Player) ShowHand() {
	for _, card := range p.Hand {
		fmt.Println(card)
	}
}
