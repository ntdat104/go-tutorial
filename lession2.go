package main

import (
	"fmt"
	"math"
	"math/bits"
)

func main() {
	// Boolean
	var myBool bool = true
	fmt.Println(myBool)

	var mySecondBool bool = false
	fmt.Println(mySecondBool)

	// String
	var myString string = "hello"
	fmt.Println(myString)

	// Integer
	var myInt int = 123
	fmt.Println(myInt)

	// int 8, 16, 32, 64
	// 1. Range
	// 2. Bits

	// 1. Range int8; -128 -> 127
	fmt.Println(math.MinInt8)
	fmt.Println(math.MaxInt8)

	// 1. Range int16; -32768 -> 32767
	fmt.Println(math.MinInt16)
	fmt.Println(math.MaxInt16)

	// 1. Range int32; -2147483648 -> 2147483647
	fmt.Println(math.MinInt32)
	fmt.Println(math.MaxInt32)

	// 1. Range int64; -9223372036854775808 -> 9223372036854775807
	fmt.Println(math.MinInt64)
	fmt.Println(math.MaxInt64)

	// 2. Bits
	fmt.Println("===========================")
	fmt.Println(bits.OnesCount8(math.MaxUint8))
	fmt.Println(bits.OnesCount16(math.MaxUint16))
	fmt.Println(bits.OnesCount32(math.MaxUint32))
	fmt.Println(bits.OnesCount64(math.MaxUint64))

	// uint - unsigned integer
	var myUInt uint = 10
	fmt.Println(myUInt)

	var myByte byte = 1
	fmt.Println(myByte)
	fmt.Printf("%T\n", myByte)

	var a byte = 'D'
	fmt.Println(a)

	// foat
	var myFloat float64 = 10.01
	fmt.Println(myFloat)

	// complex z = a + bi
	var myComplex complex64 = 10.01 + 10.01i
	fmt.Println(myComplex)

	// Rune
	var myString2 string = "Nhẫn"
	myRune:= []rune(myString2)
	for i:=0; i<len(myRune); i++ {
		fmt.Printf("%c", myRune[i])
	}
	fmt.Printf("\n")

	var a1 int = 1
	var b1 float64 = 2.1
	fmt.Println(float64(a1) + b1)

	const PI = 3.14159
	fmt.Println(PI)
}