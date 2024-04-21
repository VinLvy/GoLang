package main

import "fmt"

func clos() {
	counter := 0

	increment := func() {
		fmt.Println("increment")
		counter++
	}
	increment()
	increment()
	increment()
	fmt.Println(counter)
}
