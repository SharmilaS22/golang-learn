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

	const count = 3

	fmt.Print(var_name, var_name1, var_name2)

	// strings - immutable - have to reassign
}

func returnString(name string) string {
	return fmt.Sprintf("hello %s", name)
}
