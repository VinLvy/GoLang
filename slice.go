package main

import "fmt"

func lice() {
	names := [...]string{"daoa", "dapa", "daffa", "iyos", "doni", "gavra", "uji"}

	slice1 := names[4:6]
	fmt.Println(slice1)

	slice2 := names[:3]
	fmt.Println(slice2)

	slice3 := names[3:]
	fmt.Println(slice3)

	days := [...]string{"senin", "selasa", "rabu", "kamis", "jumat", "sabtu", "minggu"}
	dayslice1 := days[5:] //sabtu, minggu
	fmt.Println(dayslice1)

	dayslice1[0] = "sabtu baru"
	dayslice1[1] = "minggu baru"
	fmt.Println(dayslice1)
	fmt.Println(days)

	dayslice2 := append(dayslice1, "libur baru")
	fmt.Println(dayslice1)
	fmt.Println(dayslice2)
	fmt.Println(days)

	var newslice []string = make([]string, 2, 5)
	newslice[0] = "Budi"
	newslice[1] = "Budiawan"
	// newslice[2] = "Budisono" error harusnya menggunakan append

	fmt.Println(newslice)
	fmt.Println(len(newslice))
	fmt.Println(cap(newslice))

	newslice2 := append(newslice, "dafa lagi")
	fmt.Println(newslice2)
	fmt.Println(len(newslice2))
	fmt.Println(cap(newslice2))

	newslice2[0] = "dono"
	fmt.Println(newslice2)
	fmt.Println(newslice)

	fromslice := days[:]
	toslice := make([]string, len(fromslice), cap(fromslice))
	copy(toslice, fromslice)

	fmt.Println(fromslice)
	fmt.Println(toslice)

	iniarray := [...]int{1, 2, 3, 4, 5}
	inislice := []int{1, 2, 3, 4, 5}

	fmt.Println(iniarray)
	fmt.Println(inislice)
}
