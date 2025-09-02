package greatings

import (
	"errors"
	"fmt"
)

func SayHi(name string) (string, error) {
	message := fmt.Sprintf("Hi %s", name)
	if name == "" {
		return "", errors.New("name is required")
	}
	return message, nil
}
