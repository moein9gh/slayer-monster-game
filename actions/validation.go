package actions

import "errors"

func ValidateUserChoice(choice int) (err error) {
	if choice != 1 && choice != 2 && choice != 3 {
		err = errors.New("\ninvalid input")
	}
	return
}
