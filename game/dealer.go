package game

import "fmt"

type Dealer struct {
	Hand []Card
}

func (p *Dealer) AddCard(card Card) {
	p.Hand = append(p.Hand, card)
}

func (p *Dealer) ShowHand() {
	for _, card := range p.Hand {
		fmt.Println(card)
	}
}

func (p *Dealer) Play(deck *[]Card) {
	for CalculatePoints(p.Hand) < 17 {
		p.AddCard(Draw(deck))
	}
}
