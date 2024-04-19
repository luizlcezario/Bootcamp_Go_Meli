package calc

import "errors"

func Divide(a, b int) (int, error) {
	if b == 0 {
		return 0, errors.New("divide by 0 not permited")
	}
	return a / b, nil
}
