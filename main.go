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
	roundNumber = 0
	for winner == "" {
		winner = startNewRound()
	}

	endGame(winner)
}

func startGame() {
	interaction.PrintGreeting()
}

func startNewRound() (winner string) {

	healValue := 0
	damageValue := 0
	roundNumber++

	interaction.PrintRoundNumber(strconv.Itoa(roundNumber))

	isSpecialRound := roundNumber%3 == 0

	choicenAction, err := interaction.GettingSelection(isSpecialRound)

	interaction.CheckErr(err)

	userAction, err := interaction.DiagnoseAction(choicenAction)

	interaction.CheckErr(err)

	actionValue := actions.ExecutePlayerActiton(userAction)

	if userAction == "HEAL" {
		healValue = actionValue
	} else {
		damageValue = actionValue
	}

	playerHealth, monsterHealth := actions.GetHealth()

	winner = actions.CheckWinner(playerHealth, monsterHealth)
	if winner != "" {
		fmt.Println("winner is", winner)
		return
	}

	monsterAttDmg := actions.MonsterAttack()

	playerHealth, monsterHealth = actions.GetHealth()

	winner = actions.CheckWinner(playerHealth, monsterHealth)
	if winner != "" {
		fmt.Println("winner is", winner)
		return
	}

	roundData := interaction.RoundData{
		RoundNumber:         roundNumber,
		Action:              userAction,
		PlayerAttackDamage:  damageValue,
		HealValue:           healValue,
		MonsterAttackDamage: monsterAttDmg,
		MonsterHealth:       monsterHealth,
		PlayerHealth:        playerHealth,
	}

	roundData.LogRound()

	return

}

func endGame(winner string) {
	interaction.PrintWinner(winner)
}
