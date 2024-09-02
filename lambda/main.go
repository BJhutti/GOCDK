package main

import (
	"fmt"

	"github.com/aws/aws-lambda-go/lambda"
)

// structs in go can be tagged
type MyEvent struct {
	Username string `json:"username"` //unmarshalls
}

// Take in a payload and do someting with it
func HandleRequest(event MyEvent) (string, error) {
	if event.Username == "" {
		return "", fmt.Errorf("username cannot be empty")
	}

	return fmt.Sprintf("Successfully called by - %s", event.Username), nil
}

// this directory will be our backend
func main() {
	lambda.Start(HandleRequest)
}
