package greetings

import (
	"errors"
	"fmt"
)

// Hello returns a greeting for the named version
func Hello(name string) (string, error) {
	// If no name was given, return an error.
	if name == "" {
		return "", errors.New("empty name")
	}
	// If a name was received, return a value
	message := fmt.Sprintf("Hi, %v. Welcome!", name)
	return message, nil
}
