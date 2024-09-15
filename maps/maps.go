package main

import "fmt"

func main() {
	var chicken map[string]int
	chicken = map[string]int{}
	chicken["januari"] = 50
	chicken["februari"] = 40
	chicken["maret"] = 30
	fmt.Println("januari", chicken["maret"])
}
