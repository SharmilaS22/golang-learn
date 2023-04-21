package main

import "fmt"
// bool
	// string
	// float
	// int int32 int64
	// rune

func main2() {
	var var_name = 8
	// or
	var var_name1 int
	var_name1 = 10
	// or
	var_name2 := 3

	// or (type comes last)
	// var var_name, var_name1, var_name2 int

	// or
	// var var_name, var_name1, var_name2 = 8, 10, 3

	// or
	// var_name, var_name1, var_name2 := 8, 10, 3

	const count = 3 // character, string, boolean, or numeric values.
	// cannot use := for const

	fmt.Print(var_name, var_name1, var_name2)

	// zero values - 0, false, ""

	// type change type(value)
	// hello = string(2424)

	// strings - immutable - have to reassign
}

func returnString(name string) string {
	return fmt.Sprintf("hello %s", name)
}
