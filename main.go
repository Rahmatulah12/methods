package main

import (
	"fmt"
	"strings"
)

type student struct {
	name  string
	age   int
	class int
	grade int
}

func (s student) sayHello() {
	fmt.Println("Hello,", s.name)
}

func (s student) getName(i int) string {
	// strings.Split mengubah string menjadi array dengan pemisah spasi
	return strings.Split(s.name, " ")[i-1]
}

func main() {
	var s1 = student{name: "Nama Ku Bento, Rumah Real Estate", age: 27, class: 12, grade: 6}
	s1.sayHello()

	var name = s1.getName(2)
	fmt.Println(name)

	// mengakses method pointer dari objek biasa
	var s2 = student{name: "Rahmatulah Sidik", age: 27, class: 12, grade: 6}
	s2.sayHello2()

	// pengaksesan method pointer dari objek pointer
	var s3 = &student{"Siskaeee", 30, 6, 3}
	fmt.Println(&s3)
	s3.sayHello2()

	var hobbies string = "Makan,Minum,Mandi,Cebok"
	var splitHobbies = strings.Split(hobbies, ",")
	fmt.Println(splitHobbies)
}

/*
	Method Pointer
	method yang variabel objeknya dideklarasikan dalam bentuk pointer.
	Kelebihan method jenis ini adalah manipulasi data pointer
	pada property milik variabel tersebut bisa dilakukan.
*/
func (s *student) sayHello2() {
	fmt.Println("Hello,", s.name)
}
