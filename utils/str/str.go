package str

import (
	"errors"
)

func CharAt(str string, index int) (string, error) {
	if index >= len(str) {
		return "", errors.New("index exceeds string bounds")
	}

	return string(str[index]), nil
}
