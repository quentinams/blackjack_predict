package game

import (
	"awesomeProject/predict"
	"fmt"
	"strconv"
	"strings"
)

func Game(strategyTable map[string]predict.Action, number_test int) {
	game_win := 0
	var game_lose = 0
	var game_tie = 0
	totalTests := number_test
	for number_test > 0 {
		var player = Player{Cash: 100, Hand: []Card{}}
		var dealer = Dealer{Hand: []Card{}}
		var deck = NewDeck()

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
					player.Cash *= 2
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

					if dealerPoints > 21 || playerPoints1 > dealerPoints {
						game_win++
					} else if playerPoints1 < dealerPoints {
						game_lose++
					} else {
						game_tie++
					}

					if dealerPoints > 21 || playerPoints2 > dealerPoints {
						game_win++
					} else if playerPoints2 < dealerPoints {
						game_lose++
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
		} else if playerPoints < dealerPoints {
			game_lose++
		} else {
			game_tie++
		}
		number_test--

		// Update progress bar
		printProgressBar(totalTests-number_test, totalTests)
	}
	fmt.Printf("\nPlayer wins: %d\n", game_win)
	fmt.Printf("Player loses: %d\n", game_lose)
	fmt.Printf("Tie games: %d\n", game_tie)
	winRate := float64(game_win) / float64(game_win+game_lose+game_tie) * 100
	fmt.Printf("Win rate: %.2f%%\n", winRate)
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

func CalculatePoints(hand []Card) int {
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
