package main

import (
	"fmt"
	"strconv"
)

func main() {
	var input string
	fmt.Print("type some number : ")
	fmt.Scanln(&input)
	var number int
	var err error
	number, err = strconv.Atoi(input)
	if err == nil {
		fmt.Println(number, "is a number")
	} else {
		fmt.Println(input, "is not a number")
		fmt.Println(err.Error())
	}
}
