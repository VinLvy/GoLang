package main

import "fmt"

type Blacklist func(string) bool

func registeruser(name string, blacklist Blacklist) {
	if blacklist(name) {
		fmt.Println("you're blocked ", name)
	} else {
		fmt.Println("welcome", name)
	}
}
func anon() {
	blacklist := func(name string) bool {
		return name == "anjing"
	}
	registeruser("daoa", blacklist)
	registeruser("anjing", func(name string) bool {
		return name == "anjing"
	})
}
