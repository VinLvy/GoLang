package main

import "fmt"

func cons() {
	const name1 string = "Daoa"
	const name2 = "Dapa"

	fmt.Println(name1)
	fmt.Println(name2)

	//error
	//name1 = "Tidak bisa diubah"
	//name2 = "Tidak bisa diubah"
}
