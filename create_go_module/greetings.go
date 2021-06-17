package greetings

import (
	"errors"
	"fmt"
)

func Hello(name string) (string, error) {
	if name == "" { return "", errors.New("empty name") }
	message := fmt.Sprintf("Hello %v, Nice to meet you", name)
	return message, nil
}