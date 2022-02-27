package actions

import "fmt"

var currentMonsterHealth = MAX_PLAYER_HEALTH

func MonsterAttack() {
	minDmg := MIN_MONSTER_ATTACK_HEALTH
	maxDmg := MAX_MONSTER_ATTACK_HEALTH

	dmgValue := generateRandNum(minDmg, maxDmg)

	if currentPlayerHealth-dmgValue < MIN_PLAYER_HEALTH {
		currentPlayerHealth = MIN_PLAYER_HEALTH
	} else {
		currentPlayerHealth -= dmgValue
	}
	fmt.Println("sdasdasd", currentPlayerHealth, dmgValue)

}
