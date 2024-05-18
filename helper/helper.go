package helper

var version = "1.0.0" //tdk bisa diakses dari package lain
var Application = "golang"

// tdk bisa diakses dari luar package
func saygoodbye(name string) string {
	return "goodbye " + name
}

func Sayhelloo(name string) string {
	return "Hello " + name
}
