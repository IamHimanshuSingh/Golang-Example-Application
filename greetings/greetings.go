package greetings

import (
	"errors"
	"fmt"
	"math/rand"
	"time"
)

//Hello function returns the greeting for the name passed
func Hello(name string) (string, error) {
	// If no name was given, return an error with a message.
	if name == "" {
		return "", errors.New("empty name")
	}
	//Return a greeting that embeds name in the message
	message := fmt.Sprintf(randomFormat(), name)
	return message, nil
}

//Now we need to send greetings to multiple people
func Hellos(names []string) (map[string]string, error) {

	//define a map
	messages := make(map[string]string)

	//iterate over all the names and generate messages
	for _, name := range names {
		message, err := Hello(name)
		if err != nil {
			return nil, err
		}

		messages[name] = message
	}
	return messages, nil
}

func init() {
	rand.Seed(time.Now().UnixNano())
}

func randomFormat() string {
	// A Slice of message formats
	formats := []string{
		"Hi, %v. Welcome!",
		"Great to see you, %v",
		"Hail, %v! Well met!",
	}

	return formats[rand.Intn(len(formats))]
}
