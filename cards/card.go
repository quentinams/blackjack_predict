package cards

type Card struct {
	Name  string
	Value int
}

var (
	As     Card
	Deux   Card
	Trois  Card
	Quatre Card
	Cinq   Card
	Six    Card
	Sept   Card
	Huit   Card
	Neuf   Card
	Dix    Card
	Valet  Card
	Dame   Card
	Roi    Card
)

func init() {
	As = Card{Name: "As", Value: 11}
	Deux = Card{Name: "2", Value: 2}
	Trois = Card{Name: "3", Value: 3}
	Quatre = Card{Name: "4", Value: 4}
	Cinq = Card{Name: "5", Value: 5}
	Six = Card{Name: "6", Value: 6}
	Sept = Card{Name: "7", Value: 7}
	Huit = Card{Name: "8", Value: 8}
	Neuf = Card{Name: "9", Value: 9}
	Dix = Card{Name: "10", Value: 10}
	Valet = Card{Name: "Valet", Value: 10}
	Dame = Card{Name: "Dame", Value: 10}
	Roi = Card{Name: "Roi", Value: 10}
}

// GetValue retourne la valeur de la carte en tenant compte de la règle spéciale de l'As
func (c Card) GetValue(currentTotal int) int {
	if c.Name == "As" {
		// Si l'As vaut 11 ferait dépasser 21, on le compte comme 1
		if currentTotal+11 > 21 {
			return 1
		}
		return 11
	}
	return c.Value
}

// GetHandValue calcule le total d'une main de cartes
func GetHandValue(cards []Card) int {
	total := 0
	for _, card := range cards {
		total += card.GetValue(total)
	}
	return total
}

// Exemple d'utilisation :
// main := []Card{As, Six} // Total = 17 (As = 11, Six = 6)
// main2 := []Card{As, Six, Cinq} // Total = 12 (As = 1, Six = 6, Cinq = 5)
