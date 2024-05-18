package main

import "fmt"

func decl() {
	//digunakan untuk membuat alias trhdp tipe data yg sdh ada, dngn tujuan agar lbh mudah dimengerti
	type NoKTP string

	var KTPdap NoKTP = "1111111"
	var contoh string = "222222"
	var contohKTP NoKTP = NoKTP(contoh)

	fmt.Println(KTPdap)
	fmt.Println(contohKTP)
}
