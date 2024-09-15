package main

import "fmt"

func main() {
	fmt.Println("halo")
	defer fmt.Println("selamat datang")
	fmt.Println("lagi")
}
