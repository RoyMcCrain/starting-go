package main

import (
	"fmt"
	"math"
)

// Package変数 乱用禁止
var l = 100

func main() {
	var (
		x, y int
		name string
	)
	x, y = 1, 3
	name = "aaa"
	a := "bbb"
	b := 3.111
	n := one()

	fmt.Printf("1=%T,2=%v,3=%#v,4=%v, 5=%v, 6=%v\n", x, y, name, a, b, n)
	// comment
	l = l - 10
	fmt.Println(l)

	fmt.Printf("uint32 max value = %v\n", math.MaxUint32)
	doSomething(1111111111, 2)
	fmt.Printf("型%T", 1.0) // defaultはfloat64
	printFloat()
	printFloat32()
	printRune()
	printArray()
}

func one() int {
	return 1
}

func doSomething(a, b uint32) bool {
	if (math.MaxUint32 - a) < b {
		// オーバーフローするのでfalse
		return false
	} else {
		sum := a + b
		fmt.Println(sum)
	}
	return true
}

func printFloat() {
	fmt.Printf("value = %v\n", 1.0000000000000000)
	fmt.Printf("value = %v\n", 1.0000000000000001)
	fmt.Printf("value = %v\n", 1.0000000000000002)
	fmt.Printf("value = %v\n", 1.0000000000000003)
	fmt.Printf("value = %v\n", 1.0000000000000004)
	fmt.Printf("value = %v\n", 1.0000000000000005)
	fmt.Printf("value = %v\n", 1.0000000000000006)
	fmt.Printf("value = %v\n", 1.0000000000000007)
	fmt.Printf("value = %v\n", 1.0000000000000008)
	fmt.Printf("value = %v\n", 1.0000000000000009)
}

func printFloat32() {
	fmt.Printf("float32 value = %v\n", float32(1.0000000000000000))
	fmt.Printf("float32 value = %v\n", float32(1.0000000000000001))
	fmt.Printf("float32 value = %v\n", float32(1.0000000000000002))
	fmt.Printf("float32 value = %v\n", float32(1.0000000000000003))
	fmt.Printf("float32 value = %v\n", float32(1.0000000000000004))
	fmt.Printf("float32 value = %v\n", float32(1.0000000000000005))
	fmt.Printf("float32 value = %v\n", float32(1.0000000000000006))
	fmt.Printf("float32 value = %v\n", float32(1.0000000000000007))
	fmt.Printf("float32 value = %v\n", float32(1.0000000000000008))
	fmt.Printf("float32 value = %v\n", float32(1.0000000000000009))
}

func printRune() {
	r := '松'
	fmt.Printf("Rune %v, Type %T\n", r, r)
	s := "松"
	fmt.Printf("String %v, Type %T\n", s, s)
}

func printArray() {
	a := [5]int{1, 2, 3, 4, 5}
	fmt.Printf("Array %v\n", a)
	a1 := [...]int{1, 2, 3, 4, 5}
	fmt.Printf("Array %v, 1=%v\n", a1, a1[0])
}
