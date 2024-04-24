package main

import "fmt"

type Filter func(string) string

//func sayhellowithfilter(name string, filter func(string) string) {
func sayhellowithfilter(name string, filter Filter) {
	filteredname := filter(name)
	fmt.Println("hello ", filteredname)
}

func spamfilter(name string) string {
	if name == "anjing" {
		return "..."
	} else {
		return name
	}
}

func fuas() {
	sayhellowithfilter("daoa", spamfilter)

	filter := spamfilter
	sayhellowithfilter("anjing", filter)
}
