package actions

import (
	"fmt"
	"math/rand"
	"time"
)

var randSource = rand.NewSource(time.Now().UnixNano())
var currentPlayerHealth = MAX_MONSTER_HEALTH

var randGenerator = rand.New(randSource)

func ExecutePlayerActiton(userChoice string) {

	switch userChoice {

	case "ATTACK":
		PlayerAttack(false)

	case "HEAL":
		PlayerHeal()

	case "SPECIAL ATTACK":
		PlayerAttack(true)

	}

}

func PlayerAttack(specialAttack bool) {

	minDmg := MIN_PLAYER_ATTACK_HEALTH
	maxDmg := MAX_PLAYER_ATTACK_HEALTH

	if !specialAttack {
		minDmg = MIN_PLAYER_SPECIAL_ATTACK_HEALTH
		maxDmg = MAX_PLAYER_SPECIAL_ATTACK_HEALTH
	}

	dmgValue := generateRandNum(minDmg, maxDmg)

	fmt.Println(currentMonsterHealth, dmgValue)

	if currentMonsterHealth-dmgValue < MIN_MONSTER_HEALTH {
		currentMonsterHealth = MIN_MONSTER_HEALTH
	} else {
		currentMonsterHealth -= dmgValue
	}

}

func PlayerHeal() {

	healValue := generateRandNum(MIN_PLAYER_HEAL, MAX_PLAYER_HEAL)

	fmt.Println("heal", currentPlayerHealth)

	if currentPlayerHealth+healValue > MAX_PLAYER_HEALTH {
		currentPlayerHealth = MAX_PLAYER_HEALTH
	} else {
		currentPlayerHealth += healValue
	}

	fmt.Println("heal", currentPlayerHealth)
}

func generateRandNum(min int, max int) int {
	return randGenerator.Intn(max-min) + min
}

func GetHealth() (int, int) {
	return currentPlayerHealth, currentMonsterHealth
}

func CheckWinner(pHealth, mHealth int) string {

	if pHealth <= MIN_PLAYER_HEALTH {
		return "MONSTER"
	} else if mHealth <= MIN_MONSTER_HEALTH {
		return "PLAYER"
	} else {
		return ""
	}

}
