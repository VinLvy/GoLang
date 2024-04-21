package main

import "fmt"

func aray() {
	var names [4]string
	names[0] = "Daoa"
	names[1] = "Dapa"
	names[2] = "Daffa"
	names[3] = "Dappa"

	fmt.Println(names[0])
	fmt.Println(names[1])
	fmt.Println(names[2])
	fmt.Println(names[3])

	var values = [...]int{
		20, 25, 50, 100,
	}
	fmt.Println(values)
	fmt.Println(len(values))
	values[2] = 55
	fmt.Println(values)
}
