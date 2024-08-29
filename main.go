package main

import (
	"awesomeProject/input"
	"awesomeProject/predict"
	"fmt"
)

func main() {
	dealer, player := input.CheckArgs()
	fmt.Println("Dealer cards:", dealer)
	fmt.Println("Player cards:", player)

	strategyTable, err := predict.LoadStrategyTable("csv_strategie/strat.csv")
	if err != nil {
		fmt.Println("Erreur lors du chargement de la table de stratégie :", err)
		return
	}
	playerHand := player
	dealerCard := dealer
	action := predict.Predict(strategyTable, playerHand, dealerCard)
	fmt.Printf("Action recommandée pour la main %s contre la carte du croupier %s : %s\n", playerHand, dealerCard, action)
}
