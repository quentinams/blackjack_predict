package main

import (
	"awesomeProject/game"
	"awesomeProject/predict"
	"flag"
	"fmt"
	"log"
)

func main() {
	// Définition des flags
	numberTests := flag.Int("tests", 1000, "Nombre de parties à simuler")
	initialCash := flag.Int("cash", 10000, "Mise initiale par partie")
	flag.Parse()

	// Vérification des arguments
	if *numberTests <= 0 {
		log.Fatal("Le nombre de tests doit être positif")
	}
	if *initialCash <= 0 {
		log.Fatal("La mise initiale doit être positive")
	}

	// Chargement de la table de stratégie
	strategyTable, err := predict.LoadStrategyTable("csv_strategie/strat.csv")
	if err != nil {
		log.Fatal("Erreur lors du chargement de la table de stratégie :", err)
	}

	// Affichage des paramètres
	fmt.Printf("Démarrage de la simulation avec:\n")
	fmt.Printf("- Nombre de parties: %d\n", *numberTests)
	fmt.Printf("- Mise initiale: %d\n", *initialCash)
	fmt.Println("----------------------------------------")

	// Lancement de la simulation
	game.Game(strategyTable, *numberTests, *initialCash)
}
