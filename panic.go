package main

import "fmt"

func endapp() {
	fmt.Println("end app")
	message := recover()
	fmt.Println("terjadi panic", message)
}

func runapp(error bool) {
	defer endapp()
	if error {
		panic("ups error")
	}
}

func pani() {
	runapp(true)
	fmt.Println("daoa")
}
