package main

import "fmt"

type Addresss struct {
	City, Province, Country string
}

func changecountrytoindonesia(address *Addresss) {
	address.Country = "indonesia"
}

func pofu() {
	var address Addresss = Addresss{}
	changecountrytoindonesia(&address)

	fmt.Println(address)
}
