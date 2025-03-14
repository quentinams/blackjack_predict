package game

import (
	"awesomeProject/predict"
	"fmt"
	"strconv"
	"strings"
)

func Game(strategyTable map[string]predict.Action, number_test int, initialCash int) {
	game_win := 0
	game_lose := 0
	game_tie := 0
	totalTests := number_test
	totalCash := 0
	currentBet := 5
	// Création d'un seul joueur qui garde son argent entre les parties
	var player = Player{Cash: initialCash, Hand: []Card{}}
	
	for number_test > 0 {
		// Réinitialisation de la main du joueur pour chaque partie
		player.Hand = []Card{}
		var dealer = Dealer{Hand: []Card{}}
		var deck = NewDeck()

		// Vérification si le joueur a assez d'argent pour jouer
		if player.Cash - currentBet < 5 {
			fmt.Printf("\nLe joueur n'a plus assez d'argent pour continuer (Cash: %d)\n", player.Cash)
			break
		}

		player.AddCard(Draw(&deck))
		player.AddCard(Draw(&deck))
		dealer.AddCard(Draw(&deck))
		dealer.AddCard(Draw(&deck))

		var error_for int = 0
		var playerHasDoubled = false

		for {
			if error_for == 15 {
				error_for = -1
				break
			}
			error_for++
			action := predict.Predict(strategyTable, strconv.Itoa(CalculatePoints(player.Hand)), strconv.Itoa(dealer.Hand[0].Value))

			if action == "Hit" {
				player.AddCard(Draw(&deck))
				if CalculatePoints(player.Hand) > 21 {
					break
				}
			} else if action == "Stand" {
				break
			} else if action == "Double" {
				if len(player.Hand) == 2 && !playerHasDoubled {
					player.AddCard(Draw(&deck))
					playerHasDoubled = true
					currentBet *= 2
					break
				}
			} else if action == "Split" {
				if len(player.Hand) == 2 && player.Hand[0].Value == player.Hand[1].Value {
					splitHand1 := []Card{player.Hand[0], Draw(&deck)}
					splitHand2 := []Card{player.Hand[1], Draw(&deck)}

					// Play both hands
					playHand := func(hand []Card) int {
						for {
							action := predict.Predict(strategyTable, strconv.Itoa(CalculatePoints(hand)), strconv.Itoa(dealer.Hand[0].Value))
							if action == "Hit" {
								hand = append(hand, Draw(&deck))
								if CalculatePoints(hand) > 21 {
									break
								}
							} else if action == "Stand" {
								break
							} else if action == "Double" {
								if len(hand) == 2 {
									hand = append(hand, Draw(&deck))
									break
								}
							} else {
								break
							}
						}
						return CalculatePoints(hand)
					}

					playerPoints1 := playHand(splitHand1)
					playerPoints2 := playHand(splitHand2)

					dealer.Play(&deck)
					dealerPoints := CalculatePoints(dealer.Hand)

					// Gestion des gains/pertes pour la première main splitée
					if dealerPoints > 21 || playerPoints1 > dealerPoints {
						game_win++
						player.Cash += currentBet
					} else if playerPoints1 < dealerPoints {
						game_lose++
						player.Cash -= currentBet
					} else {
						game_tie++
					}

					// Gestion des gains/pertes pour la deuxième main splitée
					if dealerPoints > 21 || playerPoints2 > dealerPoints {
						game_win++
						player.Cash += currentBet
					} else if playerPoints2 < dealerPoints {
						game_lose++
						player.Cash -= currentBet
					} else {
						game_tie++
					}
					number_test--
					printProgressBar(totalTests-number_test, totalTests)
					continue
				}
			} else {
				break
			}
		}
		if error_for == -1 {
			continue
		}

		dealer.Play(&deck)
		playerPoints := CalculatePoints(player.Hand)
		dealerPoints := CalculatePoints(dealer.Hand)

		if dealerPoints > 21 || playerPoints > dealerPoints {
			game_win++
			player.Cash += currentBet
		} else if playerPoints < dealerPoints {
			game_lose++
			player.Cash -= currentBet
		} else {
			game_tie++
		}
		totalCash += player.Cash
		number_test--

		// Update progress bar
		printProgressBar(totalTests-number_test, totalTests)
	}
	fmt.Printf("\nPlayer wins: %d\n", game_win)
	fmt.Printf("Player loses: %d\n", game_lose)
	fmt.Printf("Tie games: %d\n", game_tie)
	winRate := float64(game_win) / float64(game_win+game_lose+game_tie) * 100
	fmt.Printf("Win rate: %.2f%%\n", winRate)
	
	// Calcul et affichage des statistiques d'argent
	profitLoss := float64(player.Cash) - float64(initialCash)
	profitLossPercentage := (profitLoss / float64(initialCash)) * 100
	
	fmt.Printf("\nStatistiques d'argent:\n")
	fmt.Printf("Argent final: %.2f\n", float64(player.Cash))
	fmt.Printf("Profit/Perte total: %.2f\n", profitLoss)
	fmt.Printf("Rendement: %.2f%%\n", profitLossPercentage)
}

func printProgressBar(current, total int) {
	percent := float64(current) / float64(total) * 100
	barLength := 50
	progress := int(percent / 100 * float64(barLength))
	bar := strings.Repeat("=", progress) + strings.Repeat(" ", barLength-progress)
	fmt.Printf("\r[%s] %3d%%", bar, int(percent))
	if current == total {
		fmt.Println() // Move to the next line after completion
	}
}

func ShowDeck(deck []Card) {
	for _, card := range deck {
		fmt.Println(card)
	}
}

func GetHandValue(hand []Card) int {
	points := 0
	ace := 0
	for _, card := range hand {
		points += card.Value
		if card.Rank == Ace {
			ace++
		}
	}
	for ace > 0 && points+10 <= 21 {
		points += 10
		ace--
	}
	return points
}
