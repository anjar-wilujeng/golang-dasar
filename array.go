package main

import "fmt"

func main() {
	var names [3]string

	names[0] = "Budi"
	names[1] = "Ari"
	names[2] = "Susilo"

	fmt.Println(names[0])
	fmt.Println(names[1])
	fmt.Println(names[2])

	var values = [...]int {
		90,
		80,
		30,
		100,
		80,
	}

	fmt.Println(values)
	fmt.Println(values[2])

	fmt.Println(len(values))

}