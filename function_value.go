package main

import "fmt"

func getgoodbye(name string) string {
	return "goodbye " + name
}

func fuva() {
	misal := getgoodbye
	contoh := getgoodbye
	fmt.Println(misal("daoa"))
	fmt.Println(contoh("lulo"))
}
