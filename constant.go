package main

import "fmt"

func main() {
	const firstName string = "Ali"
	const lastName = "Gugus"

	//error
	//firstName = "Joko"
	//lastName = "Budi"

	fmt.Println(firstName)
	fmt.Println(lastName)

	const (
		namaDepan string = "Joko"
		namaBelakang = "Budi"
	)
	fmt.Println(namaDepan)
	fmt.Println(namaBelakang)
}