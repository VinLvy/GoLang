package main

import "fmt"

func math() {
	var a = 5
	var b = 10
	var c = 2
	var d = a + b*c
	fmt.Println(d)

	var i = 10
	i += 10
	fmt.Println(i)
	i += 5
	fmt.Println(i)

	var j = 1
	j++
	j++
	fmt.Println(j)

}
