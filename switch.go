package main

import "fmt"

func main() {
	name := "Budi"

	switch name {
	case "Eko":
		fmt.Println("Hallo Eko")
	case "Budi":
		fmt.Println("Hallo Budi")
	default:
		fmt.Println("Anda siapa?")
	}

	switch length := len(name); length > 5 {
	case true:
		fmt.Println("nama terlalu panjang")
	case false:
		fmt.Println("nama terlalu pendek")
	}

	length := len(name)
	switch {
	case length > 5:
		fmt.Println("nama terlalu panjang")
	case length > 10:
		fmt.Println("nama terlalu panjang sekali")
	default:
		fmt.Println("nama sudah benar")
	}
}