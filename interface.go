package main

import "fmt"

type HasName interface {
	GetName() string
}

func sayHello(value HasName) {
	fmt.Println("hello", value.GetName())
}

type Person struct {
	name string
}

func (person Person) GetName() string {
	return person.name
}

type Animal struct {
	Name string
}

func (animal Animal) GetName() string {
	return animal.Name
}

func intr() {
	person := Person{name: "daoa"}
	animal := Animal{Name: "kucing"}
	sayHello(person)
	sayHello(animal)
}
