package main

import "fmt"

func foro() {
	//counter := 1

	//for counter <= 10 {
	//	fmt.Println("perulangan ke ", counter)
	//	counter++
	//}

	for counter := 1; counter <= 10; counter++ {
		fmt.Println("perulangan ke ", counter)
	}
	fmt.Println("selesai")

	names := []string{"daoa", "dapa", "daffa"}
	for i := 0; i < len(names); i++ {
		fmt.Println(names[i])
	}
	for index, name := range names {
		fmt.Println("index", index, "=", name)
	}
	for _, name := range names {
		fmt.Println(name)
	}
}
