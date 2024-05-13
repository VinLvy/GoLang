package main

import "fmt"

func factorialloop(value int) int {
	result := 1
	for i := value; i > 0; i-- {
		result *= i
	}
	return result
}

func factorialrecursive(value int) int {
	if value == 1 {
		return 1
	} else {
		return value * factorialrecursive(value-1)
	}
}

func recu() {
	result := 10 * 9 * 8 * 7 * 6 * 5 * 4 * 3 * 2 * 1
	fmt.Println(result)
	fmt.Println(factorialloop(10))
	fmt.Println(factorialrecursive(10))
}
