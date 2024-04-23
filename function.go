package main

import "fmt"

func sayhello() {
	fmt.Println("HELLO")
}

func fucc() {
	sayhello()
	sayhello()
	sayhello()
}
