package main

import "fmt"

func gethello(name string) string {
	hello := "hello" + name
	return hello
}
func retr() {
	result := gethello("daoa")
	fmt.Println(result)

	fmt.Println(gethello("daffa"))
	fmt.Println(gethello("dapa"))
}
