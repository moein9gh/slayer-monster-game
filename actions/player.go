package actions

import (
	"math/rand"
	"time"
)

var randSource = rand.NewSource(time.Now().UnixNano())
var currentPlayerHealth = MAX_MONSTER_HEALTH
var currentMonsterHealth = MAX_PLAYER_HEALTH

var randGenerator = rand.New(randSource)

func ExecutePlayerActiton(userChoice string) (val int) {

	switch userChoice {

	case "ATTACK":
		val = PlayerAttack(false)

	case "HEAL":
		val = PlayerHeal()

	case "SPECIAL ATTACK":
		val = PlayerAttack(true)
	}

	return
}

func PlayerAttack(specialAttack bool) int {

	minDmg := MIN_PLAYER_ATTACK_HEALTH
	maxDmg := MAX_PLAYER_ATTACK_HEALTH

	if specialAttack {
		minDmg = MIN_PLAYER_SPECIAL_ATTACK_HEALTH
		maxDmg = MAX_PLAYER_SPECIAL_ATTACK_HEALTH
	}

	dmgValue := generateRandNum(minDmg, maxDmg)

	if currentMonsterHealth-dmgValue < MIN_MONSTER_HEALTH {
		currentMonsterHealth = MIN_MONSTER_HEALTH
	} else {
		currentMonsterHealth -= dmgValue
	}

	return dmgValue
}

func MonsterAttack() int {
	minDmg := MIN_MONSTER_ATTACK_HEALTH
	maxDmg := MAX_MONSTER_ATTACK_HEALTH

	dmgValue := generateRandNum(minDmg, maxDmg)

	if currentPlayerHealth-dmgValue < MIN_PLAYER_HEALTH {
		currentPlayerHealth = MIN_PLAYER_HEALTH
	} else {
		currentPlayerHealth -= dmgValue
	}

	return dmgValue

}

func PlayerHeal() int {

	healValue := generateRandNum(MIN_PLAYER_HEAL, MAX_PLAYER_HEAL)

	if currentPlayerHealth+healValue > MAX_PLAYER_HEALTH {
		currentPlayerHealth = MAX_PLAYER_HEALTH
	} else {
		currentPlayerHealth += healValue
	}

	return healValue
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
