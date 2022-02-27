package interaction

import (
	"errors"
	"fmt"

	"github.com/moein9gh/slayer-monster-game/actions"
)

func GettingSelection(isSpecialRound bool) (playerAction int, err error) {

	emptySelection := -1

	for emptySelection == -1 {

		playerAction = 0

		ShowOptions(isSpecialRound)
		_, err = fmt.Scan(&playerAction)

		CheckErr(err)
		err = actions.ValidateUserChoice(playerAction)

		if err == nil {
			break
		} else {
			PrintInvalidData()
		}
	}
	return
}

func DiagnoseAction(actionNumber int) (action string, err error) {

	switch actionNumber {
	case 1:
		action = "ATTACK"
	case 2:
		action = "HEAL"
	case 3:
		action = "SPECIAL ATTACK"
	default:
		err = errors.New("\ninvalid input")
		return
	}

	return

}
