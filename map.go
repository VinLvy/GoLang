package main

import "fmt"

func mapp() {
	//var person map[string]string = map[string]string{}
	//person["name"] = "daoa"
	//person["address"] = "kedungwaru"

	person := map[string]string{
		"name":    "dapa",
		"address": "kedungwaru",
	}
	fmt.Println(person["name"])
	fmt.Println(person["address"])
	fmt.Println(person)

	book := make(map[string]string)
	book["title"] = "buku saya bundar"
	book["author"] = "daoa"
	book["wrong"] = "ups"

	fmt.Println(book)
	delete(book, "wrong")
	fmt.Println(book)
}
