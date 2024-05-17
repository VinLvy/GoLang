package main

import "fmt"

func witch() {
	name := "dappa"

	switch name {
	case "daoa":
		fmt.Println("hello daoa")
	case "dappa":
		fmt.Println("hello dappa")
	case "dafa":
		fmt.Println("hello dafa")
	default:
		fmt.Println("nama kamu siapa?")
	}

	switch length := len(name); length > 5 {
	case true:
		fmt.Println("nama terlalu panjang")
	case false:
		fmt.Println("nama sudah bnar")
	}

	nama := "kurniawan"
	length2 := len(nama)
	switch {
	case length2 > 10:
		fmt.Println("nama terlalu panjang")
	case length2 > 5:
		fmt.Println("nama cukup panjang")
	default:
		fmt.Println("nama sudah beenar")
	}
}
