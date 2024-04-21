package main

import "fmt"

type AAddress struct {
	City, Province, Country string
}

func aste() {
	var address1 AAddress = AAddress{"subang", "jawa barat", "indonesia"}
	var address2 *AAddress = &address1 //pointer

	address2.City = "bandung"
	fmt.Println(address1) // ikut berubah
	fmt.Println(address2) // berubah mnjd bandung

	//address2 = &AAddress{"jakarta", "DKI jakarta", "indonesia"} tdk berubah
	*address2 = AAddress{"jakarta", "DKI jakarta", "indonesia"}
	fmt.Println(address1) // berubah
	fmt.Println(address2)
}
