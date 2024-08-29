package predict

import "fmt"

func Predict(strategyTable map[string]Action, playerHand string, dealerCard string) string {
	return string(strategyTable[fmt.Sprintf("%s,%s", playerHand, dealerCard)])
}
