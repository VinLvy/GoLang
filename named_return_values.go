package main

import "fmt"

func getcompletename() (firstname, middlename, lastname string) {
	firstname = "daoa"
	middlename = "rendra"
	lastname = "anggara"
	return firstname, middlename, lastname
}

func named() {
	a, b, c := getcompletename()
	fmt.Println(a, b, c)
}
