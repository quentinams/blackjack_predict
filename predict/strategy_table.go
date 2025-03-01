package predict

import (
	"encoding/csv"
	"fmt"
	"os"
)

func actionFromCode(code string) Action {
	switch code {
	case "1":
		return Hit
	case "2":
		return Stand
	case "3":
		return Double
	case "4":
		return Split
	default:
		return Hit
	}
}

func LoadStrategyTable(filename string) (map[string]Action, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			fmt.Println("Error closing file:", err)
		}
	}(file)

	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		return nil, err
	}

	strategyTable := make(map[string]Action)

	headers := records[0][1:] // En-têtes représentant les cartes du croupier

	for _, record := range records[1:] {
		playerHand := record[0] // Main du joueur
		for i, actionCode := range record[1:] {
			dealerCard := headers[i]
			key := fmt.Sprintf("%s,%s", playerHand, dealerCard)
			strategyTable[key] = actionFromCode(actionCode)
		}
	}

	return strategyTable, nil
}
