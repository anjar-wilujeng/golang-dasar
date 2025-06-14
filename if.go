package main

import "fmt"

func main() {
	name := "Adi"

	if name == "Budi" {
		fmt.Println("Hallo Budi")
	} else if name == "Joko" {
		fmt.Println("Hallo Joko")
	} else if name == "Adi" {
		fmt.Println("Hallo Adi")
	} else {
		fmt.Println("Anda siapa?")
	}

	if length := len(name); length > 5 {
		fmt.Println("Nama sudah benar")
	} else {
		fmt.Println("Nama kurang kengkap")
	}
}