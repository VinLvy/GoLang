package main

import "fmt"

type address struct {
	City, Province, Country string
}

func poin() {
	/*
		address1 := address{"subang", "jawa barat", "indonesia"}
		address2 := address1 // copy value

		address2.City = "bandung"
		fmt.Println(address1) // tdk berubah
		fmt.Println(address2) // berubah mnjd bandung
	*/

	var address1 = address{"subang", "jawa barat", "indonesia"}
	var address2 = &address1 //pointer

	address2.City = "bandung"
	fmt.Println(address1) // ikut berubah
	fmt.Println(address2) // berubah mnjd bandung
}
