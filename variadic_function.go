package main

import "fmt"

func sumall(numbers ...int) int {
	total := 0
	for _, number := range numbers {
		total += number
	}
	return total
}

func vari() {
	//sumall(10, 10, 10)
	fmt.Println(sumall(10, 10, 10, 10, 10))

	numbers := []int{10, 10, 10}
	fmt.Println(sumall(numbers...))
}
