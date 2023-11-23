package letmein

import (
	"errors"
	"fmt"
)

func IsAuthorized(user, action, token) (bool, error) {
	err := CallLetMeIn(user, action, token)
	if errors.Is(err, ErrUnexpectedBehavior) {
		return false, fmt.Errorf("calling LET-ME-IN: %v", err)
	} else if err == ErrForbidden {
		return false, errors.New("got forbidden from LET-ME-IN")
	} else if err != nil {
		return false, err
	}

	return true, nil
}

func IsAuthorizedV2() {

}
