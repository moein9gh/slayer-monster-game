package interaction

import (
	"fmt"
	"log"

	"github.com/common-nighthawk/go-figure"
)

type RoundData struct {
	RoundNumber         int
	Action              string
	PlayerAttackDamage  int
	HealValue           int
	MonsterAttackDamage int
	MonsterHealth       int
	PlayerHealth        int
}

func PrintGreeting() {
	myFigure := figure.NewColorFigure(" MONSTER   SLAYER   GAME ", "", "red", true)
	myFigure.Print()
}

func PrintRoundNumber(roundNumber string) {
	fmt.Printf("**************************  ROUND %v *******************************\n", roundNumber)
}

func ShowOptions(isSpecialAttackAvailable bool) {
	fmt.Println("********************** Please make a choice **********************")
	fmt.Println("------------------------------------------------------------------")
	fmt.Println("1. Attack Monster")
	fmt.Println("2. Heal")
	if isSpecialAttackAvailable {
		fmt.Println("3. Special Attack")
	}

}

func CheckErr(err error) {
	if err != nil {
		log.Fatal("ERROR", err)
	}
}

func PrintInvalidData() {
	fmt.Println("********************** INVALID DATA ENTERED **********************")
}

func PrintWinner(winner string) {
	myFigure := figure.NewColorFigure(winner+"  WON! ", "", "white", true)
	myFigure.Print()
}

func PrintHealth(pHealth, mHealth int) {
	fmt.Printf("\n************ YOUR DAMAGE : %v /////// MONSTER HEALTH : %v ***********\n", pHealth, mHealth)
}

func (r *RoundData) LogRound() {
	fmt.Println("***************************** ROUND LOG *****************************")
	fmt.Printf("\nROUND:%v //// PLAYER_ACTION=%v //// ACTION_DAMAGE=%v //// HEAL_VALUE=%v //// MONSTER_DAMAGE=%v //// PLAYER_HEALTH=%v //// MONSTER_DAMAGE=%v\n", r.RoundNumber, r.Action, r.PlayerAttackDamage, r.HealValue, r.MonsterAttackDamage, r.PlayerHealth, r.MonsterHealth)
}
