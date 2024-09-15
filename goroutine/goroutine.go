package main

import (
	"fmt"
	"runtime"
)

func Print(till int, message string) {
	for i := 0; i < till; i++ {
		fmt.Println((i + 1), message)
	}
}

func main() {
	runtime.GOMAXPROCS(2)

	Print(5, "halo")
	Print(5, "apa kabar")
	var input string
	fmt.Scanln(&input)
}
