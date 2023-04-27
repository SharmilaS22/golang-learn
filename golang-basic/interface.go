package main

import "fmt"

type Name struct {
	firstName string
	lastName string
}

type printStrings interface {
	joinAndPrint()
}

// type NAme is implementing printStrings interface
func (n1 *Name) joinAndPrint () {
	fmt.Printf("Name: %s %s\n", n1.firstName, n1.lastName)
}

type Rectangle struct {
	l int
	b int
}

type Square struct {
	a int
}

type Shape interface {
	area() int
}

func (r Rectangle) area () int {
	return r.l * r.b
}

func (s *Square) area () int {
	if s == nil { 
		fmt.Print("Hello") 
		return 0 
	} //handling nil type
	fmt.Print(s)
	return 4 * s.a
}

func printArea (sh Shape) {
	fmt.Println(sh.area())
}

func main () {

	// ex 1
	n2 := Name{"Cinco", "Hargreeves"}
	n2.joinAndPrint()

	n3 := Name{"Avery", "Kylie Grambs"}
	n3.joinAndPrint()

	// ex 2
	s1 := Square{2}
	r1 := Rectangle{3, 2}

	printArea(&s1) // *Square type - need to give pointer
	printArea(r1) //Rectangle type - no need address
	
	var shape1 Shape = &s1 //
	// interface => (value, type) => (&{2}, *Square)
	fmt.Println(shape1.area())

	fmt.Println(s1, r1, shape1)

	var s2 *Square //s2 is nil
	// if s2 is not declared as pointer(*Square) -> it would have been initialised with {0}
	fmt.Println(s2.area()) // does not return error - can handle in function

	// https://go.dev/tour/methods/13

	sq := shape1.(Shape)
	// if shape1 is not Shape type then throws error (panic)
	// error handler ↓
	val, ok := shape1.(Shape) //ok = isShape or not
	// val is shape1 assigned, if ok=false => val is zero value of Type (as in val.(Type))
	fmt.Println(sq, val, ok)

	printType(1)
	printType("hello")
}

func printType(val interface{}) { 
	switch t :=  val.(type) {
	case int: fmt.Println("integer ", t)
	case string: fmt.Println("string ", t)
	default: fmt.Println("unknown type ", t)
	}
}