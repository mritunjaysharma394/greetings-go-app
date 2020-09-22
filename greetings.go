package greetings

import (
	"errors"
	"fmt"
	"math/rand"
	"time"
)

// Hello returns a greeting for the named version
func Hello(name string) (string, error) {
	// If no name was given, return an error.
	if name == "" {
		return "", errors.New("empty name")
	}
	// If a name was received, return a value
	message := fmt.Sprintf(randomFormat(), name)
	return message, nil
}

// init sets inital values for variables used in the function
func init() {
	rand.Seed(time.Now().UnixNano())
}

// randomFormat returns one of the set greetings selected at random.
func randomFormat() string {
	// A slice of message formats.
	formats := []string{
		"Hi, %v. Welcome!",
		"Great to see you, %v!",
		"Hail, %v! Well met",
	}

	//Return a randomly selected message format
	return formats[rand.Intn(len(formats))]
}
