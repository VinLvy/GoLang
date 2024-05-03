package main

import "fmt"

/*
func contoh(name string)string{
	if name == ""{
		return nil
	} else {
		return name
	}
}
*/

func newmap(name string) map[string]string {
	if name == "" {
		return nil
	} else {
		return map[string]string{
			"name": name,
		}
	}
}

func nili() {
	data := newmap("")
	if data == nil {
		fmt.Println("data masih kosong")
	} else {
		fmt.Println(data["name"])
	}
}
