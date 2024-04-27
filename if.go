package main

import "fmt"

func ifif() {
	name := "daoa"

	if name == "daoa" {
		fmt.Println("hello daoa")
	} else if name == "dono" {
		fmt.Println("hello dono")
	} else if name == "jono" {
		fmt.Println("hello jono")
	} else {
		fmt.Println("halo nama kamu siapa?")
	}

	//if length := len(name); length > 5
	length := len(name)
	if length > 5 {
		fmt.Println("nama terlalu panjang")
	} else {
		fmt.Println("nama sudah benar")
	}
}
