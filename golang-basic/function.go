package main

import (
	"fmt"
)

// type comes after the argument
// return type ---------↓---
func add(a int, b int) int {
	return a + b
}

// if two or more *consecutive* arguments with same type
func add1(a, b int) int {
	return a + b
}

// returning multiple values
func addAndMul (x, y int) (int, int) {
	return x + y, x * y
}

// naming what to return along with type
func addAndMul2(x, y int) (sum, prod int) { //the sum and prod are declared - just can assign values
	sum  = x + y
	prod = x * y
	return //naked returns - only to be used on short functions
	// affects readability in long functions
}

func main () {
	sum, prod := addAndMul2(3, 2)
	fmt.Println(sum, prod)
}

// switch does not need break statements

func switchexample(number int) string {

	switch number % 2 == 0 {
	case true: return "even"
	case false: return "odd"
	default: return "Not defined"
	}
}

// no while or do while - use for with no condition
// if else no need paranthesis

func forloop(number int) {
	sum := 0
	for i := 1; i < number; i++ {
		sum = sum + i
	}

	i := 0
	for ; ; {	
		fmt.Println("hello")
		if i == number + 3 {break}
		i++
	}
}


// defer
// run this before returning anything
// any connection closing, any resource - cleaning
//   ↓
// runs even if theres any error

func someFunc() {
	defer fmt.Println("hello!!")
	// some statements
}

