package main

import "fmt"

func main() {
	person := map[string]string{
		"name" : "Budi",
		"address" : "Jogja",
	}

	fmt.Println(person["name"])
	fmt.Println(person["address"])
	fmt.Println(person)

	book := make(map[string]string)
	book["title"] = "belajar golang"
	book["author"] = "Budi"
	book["ups"] = "Salah"

	fmt.Println(book)

	delete(book, "ups")
	fmt.Println(book)
}