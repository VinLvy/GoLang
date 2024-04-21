package main

import "fmt"

func ups() any {
	//return 1
	//return true
	return "ups"
}

func anya() {
	var kosong any = ups()
	fmt.Println(kosong)
}
