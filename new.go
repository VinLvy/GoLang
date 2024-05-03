package main

import "fmt"

type Address struct {
	City, Province, Country string
}

func neww() {
	alamat1 := new(Address) //&Address{}
	alamat2 := alamat1

	alamat2.Country = "indonesia"

	fmt.Println(alamat1)
	fmt.Println(alamat2)
}
