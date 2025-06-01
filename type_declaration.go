package main

import "fmt"

func main() {
	type NoKTP string

	var contohKTP NoKTP = "1111111111111"

	var contoh string = "2222222222222"
	var contohKTP1 NoKTP = NoKTP(contoh)

	fmt.Println(contohKTP)
	fmt.Println(contohKTP1)
}