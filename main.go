package main

import (
	"fmt"
	"strconv"

	"github.com/moein9gh/slayer-monster-game/interaction"
)

var roundNumber int = 0

func main() {
	// winner := ""
	startGame()
	x := 1
	for x == 1 {
		startNewRound()
		x++
	}
}

func startGame() {
	interaction.PrintGreeting()
}

func startNewRound() {
	roundNumber++
	interaction.PrintRoundNumber(strconv.Itoa(roundNumber))

	isSpecialRound := roundNumber%3 == 0

	interaction.ShowOptions(isSpecialRound)

	choicenAction, err := interaction.GettingSelection(isSpecialRound)

	interaction.CheckErr(err)

	userAction, err := interaction.DiagnoseAction(choicenAction)

	interaction.CheckErr(err)

	fmt.Println(userAction)

}
