package main

import "fmt"

func main() {
	// if else
	number := 10

	// if condition { // code here }
	if number == 10 {
		fmt.Println("number = 10")
	}

	// if condition { // code here } else { // code here }
	if number < 10 {
		fmt.Println("number < 10")
	} else {
		fmt.Println("number > 10")
	}

	// if statement; condition { // code here } else if condition { // code here } else { // code here }
	if a := 100; a == 100 {
		fmt.Println("a = 100")
	} else if a > 100 {
		fmt.Println("a > 100")
	} else {
		fmt.Println("a < 100")
	}

	// switch - case
	number1 := 10
	switch number1 {
		case 200,300,400:
			fmt.Println("number1 = 200,300,400")
		case 1:
			fmt.Println("number1 = 1")
		case 2:
			fmt.Println("number1 = 2")
		case 3:
			fmt.Println("number1 = 3")
		case 10:
			fmt.Println("number1 = 10")
		default:
			fmt.Println("number1 = unknown")
	}

	switch {
		case number1 > 10:
			fmt.Println("number1 > 10")
		case number1 == 10:
			fmt.Println("number1 = 10")
	}

	// Fallthrough
	number2 := 10
	switch number2 {
		case 200,300,400:
			fmt.Println("number1 = 200,300,400")
			fallthrough
		case 1:
			fmt.Println("number1 = 1")
			fallthrough
		case 2:
			fmt.Println("number1 = 2")
			fallthrough
		case 3:
			fmt.Println("number1 = 3")
			fallthrough
		case 10:
			fmt.Println("number1 = 10")
			fallthrough
		default:
			fmt.Println("number1 = unknown")
	}

	// loop
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}

	// while
	i := 0
	for i < 10 {
		fmt.Println(i)
		i++
	}
}