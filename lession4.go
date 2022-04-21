package main

import "fmt"

// func func_name (params) (return_type) { //code }

func hello() {
	fmt.Println("Hello World")
}

func hello2(name string) {
	fmt.Println("Hello", name)
}

func greeting(name string) string {
	return "Hello " + name
}

// Multiple return values
func rectInfo(width int, height int) (int, int) {
	return width, height
}

// Named return values
func rectInfo2(w int, h int) (width int, height int) {
	return w, h
}


func main() {
	hello2("Dat")
	fmt.Println(greeting("Dũng"))

	// Multiple return values
	width, height := rectInfo(10, 20)
	fmt.Println(width, height)

	// Named return values
	w, h := rectInfo2(10, 20)
	fmt.Println(w, h)
}