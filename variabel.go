package main

import "fmt"

func main() {
	// Variabel

	var firstName string = "my"
	const lastName string = "skill"

	firstName = "aku"

	// var bilanganBulat uint8 = 20

	// var bilanganDesimal = 2.0

	// var varBool = true

	var numberA int = 7
	var numberB *int = &numberA
	*numberB = 9
	fmt.Println("halo ", firstName, lastName, numberA, "!\n")
}
