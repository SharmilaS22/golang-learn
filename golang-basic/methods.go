package main

import "fmt"

type Book struct {
	title string
	author string
}

func (b1 Book) displayDetails () {
	fmt.Println(b1.title, b1.author)
}

func main () {
	b1 := Book{"AGGGTM", "Holly J"}
	b1.displayDetails()
}

// every struct is an implementer of empty interface -> so can use as a tenmplate
func print(val ...interface{}) { 
	// send each val as one by one -> add ... at end
	// if val were int[] we can just print as array
	fmt.Println(val...)
}