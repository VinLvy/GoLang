package main

import "fmt"

type Customer struct {
	name, address string
	age           int
}

func (customer Customer) sayhello(Name string) {
	fmt.Println("hello", Name, "my name is", customer.name)
}

func stru() {
	var daoa Customer
	fmt.Println(daoa)

	daoa.name = "dappa"
	daoa.address = "indonesia"
	daoa.age = 100
	fmt.Println(daoa)
	fmt.Println(daoa.name)
	fmt.Println(daoa.address)
	fmt.Println(daoa.age)

	jono := Customer{
		name:    "jonoo",
		address: "jawa",
		age:     200,
	}
	fmt.Println(jono)

	budi := Customer{"budii", "usa", 30}
	fmt.Println(budi)

	daoa.sayhello("ageng")
	jono.sayhello("ageng")
	budi.sayhello("ageng")
}
