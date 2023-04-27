package main

import "fmt"

type Book struct {
	title string
	author string
}
// method of Book
func (b1 Book) displayDetails () {
	fmt.Println(b1.title, b1.author)
}

// methods for custom types
type myInt int

// increment by one - copy
func (i myInt) increment () myInt { // pass by value
	i += 1
	return i
}
// increment in place
func (i *myInt) incrementInPlace () myInt {
	*i += 1
	return *i
}

// by default, all methods should be pointer (*Book instead of Book)
// saves memory - no saving again in new memory
func (b1 *Book) updateAuthor(newAuthor string) {
	b1.author = newAuthor
}
func main () {
	b1 := Book{"AGGGTM", "Holly J"}
	b1.displayDetails()

	var v myInt = 3
	print(v.increment())
	print(v) // changes not reflected

	v.incrementInPlace()
	print(v) //changes reflected

	book1 := Book{ "The inheritance Games", "Jennifer L B"}
	print(book1)
	book1.updateAuthor("Jennifer Lynn Barnes")
	print(book1)

	book1.displayDetails()
	book1.updateAuthor("Jennifer")
}

// every struct(custom type, or just any type) is an implementer of empty interface -> so can use as a tenmplate
func print(val ...interface{}) { 
	// send each val as one by one -> add ... at end
	// if val were int[] we can just print as array
	fmt.Println(val...)
}