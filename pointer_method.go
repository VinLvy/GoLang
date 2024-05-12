package main

import "fmt"

type Man struct {
	Name string
}

func (man *Man) Married() {
	man.Name = "Mr. " + man.Name
}

func pome() {
	daoa := Man{"Daoa"}
	daoa.Married()

	fmt.Println(daoa.Name)
}
