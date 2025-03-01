package main

import (
	"awesomeProject/game"
	"awesomeProject/predict"
	"flag"
	"fmt"
	"strconv"
)

func main() {
	fileName := flag.String("file", "csv_strategie/strat.csv", "Nom du fichier de stratégie CSV")
	numTests := flag.String("tests", "1000000", "Nombre d'essais")
	flag.Parse()
	numTestsInt, err := strconv.Atoi(*numTests)
	if err != nil {
		fmt.Println("Erreur lors de la conversion du nombre de tests :", err)
		return
	}
	strategyTable, err := predict.LoadStrategyTable(*fileName)
	if err != nil {
		fmt.Println("Erreur lors du chargement de la table de stratégie :", err)
		return
	}

	fmt.Printf("Start test of predict with %d tests\n\n", numTestsInt)
	game.Game(strategyTable, numTestsInt)
}
