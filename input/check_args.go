package input

import (
	"flag"
)

func CheckArgs() (dealer string, player string) {
	dealerArg := flag.String("dealer", "", "List of dealer cards, e.g., \"A,1,J\"")
	playerArg := flag.String("player", "", "List of player cards, e.g., \"K,10,2\"")

	flag.Parse()

	return *dealerArg, *playerArg
}
