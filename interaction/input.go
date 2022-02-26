package interaction

import (
	"errors"
	"fmt"
)

func GettingSelection(isSpecialRound bool) (playerAction int, err error) {

	playerAction = 0

	_, err = fmt.Scan(&playerAction)

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
