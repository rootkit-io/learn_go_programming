package main

import (
	"errors"
	"fmt"
)

func main() {
	err := returnError()
	if err != nil {
		fmt.Println(err)
	}
}

func returnError() error {
	return errors.New("error here")
}
