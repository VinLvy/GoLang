package main

import "fmt"

func random() any {
	return "ok"
}

func asse() {
	var result any = random()
	/*
		var resultString string = result.(string)
		fmt.Println(resultString)
		var resultInt = result.(int)
		fmt.Println(resultInt)
	*/
	switch value := result.(type) {
	case string:
		fmt.Println("string", value)
	case int:
		fmt.Println("int", value)
	default:
		fmt.Println("unknown", value)
	}
}
