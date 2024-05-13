package main

import "fmt"

func getfullname() (string, string) {
	return "daoa", "rendra"
}

func mult() {
	firstname, lastname := getfullname()
	fmt.Println(firstname, lastname)

	//firstname, _ := getfullname()
	//fmt.Println(firstname) jika tidak digunakan
}
