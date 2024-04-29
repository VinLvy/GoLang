package main

import (
	"GoLang-Lesson/database"
	_ "GoLang-Lesson/internal"

	//"GoLang-Lesson/internal" tdk bisa berjalan
	"fmt"
)

func inii() {
	fmt.Println(database.Getdatabase())
}
