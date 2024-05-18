package main

import "fmt"

func variable() {
	var names string
	names = "Rnames"
	fmt.Println(names)

	var name = "Daoa"
	fmt.Println(name)
	name = "Dapa"
	fmt.Println(name)

	name2 := "Dapin"
	fmt.Println(name2)

	var (
		firstname = "Dappa"
		lastname  = "Anggara"
	)
	fmt.Println(firstname)
	fmt.Println(lastname)
}
