package main

import (
	"fmt"

	"github.com/joloong/GoExercises/blackjack_ai/blackjack"
)

func main() {
	game := blackjack.New()
	winnings := game.Play(blackjack.HumanAI())
	fmt.Println("You have:", winnings)
}
