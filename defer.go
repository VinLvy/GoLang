package main

import "fmt"

func logging() {
	fmt.Println("selesai memanggil function")
}

func runapplication() {
	defer logging()
	fmt.Println("run application")
}

func defe() {
	runapplication()
}
