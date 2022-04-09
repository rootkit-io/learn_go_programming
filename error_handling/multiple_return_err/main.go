package main

import (
	"fmt"
	"strconv"
)

func main() {
	result, err := returnError(true)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(result)
	}
}

type specialError struct {
	errorMessage string
	errorCode	 int
}

func (d specialError) Error() string {
	return d.errorMessage + " " +strconv.Itoa(d.errorCode)
}

func returnError(returnError bool)(string, error){
	if returnError {
		return "", specialError{"Special Error", 123}
	}
	return "Random result", nil
}
