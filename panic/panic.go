package main

import (
	"errors"
	"fmt"
	"strings"
)

func validate(input string) (bool, error) {
	if strings.TrimSpace(input) == "" {
		return false, errors.New("cannot be empty")
	}
	return true, nil
}

func main() {
	var name string
	fmt.Print("enter your name: ")
	fmt.Scanln(&name)
	if valid, err := validate(name); valid {
		fmt.Println("halo", name)
	} else {
		defer fmt.Println("kebaca")
		panic(err.Error())
		fmt.Println(err.Error())
	}
}
