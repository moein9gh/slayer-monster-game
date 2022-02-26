package interaction

import (
	"fmt"
	"log"
)

func PrintGreeting() {
	fmt.Println("***************** WELCOME TO SLAYER MONSTER GAME *****************")
	fmt.Println("************************* GOOD LOCK ******************************")
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
