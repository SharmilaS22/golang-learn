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

// passing functions as values
func operations (oper func (int, int) int, x, y int) int {
	return oper(x, y)
}

// functions with closures
func adder () func (int) int {
	// returns a func (which takes an int and returns an int)
	sum := 0 // the below func is "bound" to sum variable
	return func (a int) int {
		sum += a
		return sum
	}

}

func fibonacci() func() int {
	a := 0
	b := 1
	return func () int {
		temp := a
		a = b
		b = temp + b
		return temp
	}
}

func main () {
	sum, prod := addAndMul2(3, 2)
	fmt.Println(sum, prod)


	// passing functions as values
	addFunc := func (a, b int) int {
		return a + b
	}
	mulFunc := func(a, b int) int {
		return a * b
	}
	fmt.Println(operations(addFunc, 4, 5))
	fmt.Println(operations(mulFunc, 4, 5))

// functions with closures
// A closure is a function value that references variables from outside its body.
	val := adder()
	fmt.Println(val(2)) //2
	fmt.Println(val(3)) //5

	// another example for closure
	f := fibonacci() //returns func
	for i := 0; i < 5; i++ {
		fmt.Println(f())
	}

}

// switch does not need break statements
// switch w/o condition -> if-then-else chains
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

	for i < number {	
		fmt.Println("hello")
		i++
	}

	for {	
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

func funcWithDefer() {
	defer fmt.Println("world")
	fmt.Println("hello")
	//o/p
	// hello world
}
func anotherDefer() {
	fmt.Println("counting")

	for i := 0; i < 5; i++ {
		defer fmt.Println(i) //added to stack LIFO
	}

	fmt.Println("done")
	// o/p
	// counting done 4 3 2 1
}

func ifcond() {
	if num := len("hello"); num == 5 { // num - scope only if and else
		fmt.Println("can declare variables and can use it like this in if condition")
	} else {
		fmt.Println(num) //can access variables declared in if
	}
	// cant use 'num' here (outside if-else blocks)
}





