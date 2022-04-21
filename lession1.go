package main

import "fmt"

func main() {
	var number int
	number = 10
	fmt.Println(number)

	var number1 int = 11
	fmt.Println(number1)

	// Type inference
	var email = "ntd.k108cva@gmail.com";
	fmt.Println(email)

	// Khai báo nhiều biến
	// 1. Khai báo nhiều biến cùng kiểu dữ liệu
	// 2. Khai báo nhiều biến nhưng khác nhau kiểu dữ liệu

	// Cùng kiểu dữ liệu
	var a, b int
	a = 1
	b = 2
	fmt.Println(a)
	fmt.Println(b)

	var a1, b1 int = 10, 11
	fmt.Println(a1)
	fmt.Println(b1)

	var a2, b2 = 10, 11
	fmt.Println(a2)
	fmt.Println(b2)

	//  Khác kiểu dữ liệu
	var (
		name string
		address string
		age int
	)
	name = "Nguyen Thi Dung"
	address = "Hai Chau"
	age = 20
	fmt.Println(name)
	fmt.Println(address)
	fmt.Println(age)

	var (
		name1 string = "Nguyen Thi Dung"
		address1 string = "Hai Chau"
		age1 int = 20
	)
	fmt.Println(name1)
	fmt.Println(address1)
	fmt.Println(age1)

	var name2, address2, age2 = "Nguyen Thi Dung", "Hai Chau", 20
	fmt.Println(name2)
	fmt.Println(address2)
	fmt.Println(age2)

	language := "Go"
	fmt.Println(language)
}