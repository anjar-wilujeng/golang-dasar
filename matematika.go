package main

import "fmt"

func main() {
	var a = 20
	var b = 10
	var c = a + b
	var d = a - b
	var e = a * b
	var f = a / b
	var g = a % b

	fmt.Println(c)
	fmt.Println(d)
	fmt.Println(e)
	fmt.Println(f)
	fmt.Println(g)

	var i = 10
	i += 10
	fmt.Println(i)

	i -= 10
	fmt.Println(i)

	i *= 10
	fmt.Println(i)

	i /= 10
	fmt.Println(i)


	var j = 1
	j++ //j = j + 1
	fmt.Println(j)
	j++
	fmt.Println(j)

	j-- //j = j - 1
	fmt.Println(j)
	j--
	fmt.Println(j)
}