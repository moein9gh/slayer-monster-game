package main

import (
	"fmt"
	"strconv"

	"github.com/moein9gh/slayer-monster-game/actions"
	"github.com/moein9gh/slayer-monster-game/interaction"
)

var roundNumber int = 0

func main() {
	winner := ""
	startGame()
	x := 1
	for winner == "" {
		winner = startNewRound()
		x++
	}

	endGame(winner)
}

func startGame() {
	interaction.PrintGreeting()
}

func startNewRound() (winner string) {
	roundNumber++
	interaction.PrintRoundNumber(strconv.Itoa(roundNumber))

	isSpecialRound := roundNumber%3 == 0

	choicenAction, err := interaction.GettingSelection(isSpecialRound)

	interaction.CheckErr(err)

	userAction, err := interaction.DiagnoseAction(choicenAction)

	interaction.CheckErr(err)

	actions.ExecutePlayerActiton(userAction)

	playerHealth, monsterHealth := actions.GetHealth()

	winner = actions.CheckWinner(playerHealth, monsterHealth)
	if winner != "" {
		fmt.Println("winner is", winner)
		return
	}

	actions.MonsterAttack()

	playerHealth, monsterHealth = actions.GetHealth()

	winner = actions.CheckWinner(playerHealth, monsterHealth)
	if winner != "" {
		fmt.Println("winner is", winner)
		return
	}

	interaction.PrintHealth(playerHealth, monsterHealth)

	return

}

func endGame(winner string) {
	interaction.PrintWinner(winner)
}
