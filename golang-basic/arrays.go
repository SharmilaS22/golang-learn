package main

import (
	"fmt"
)


func main3() {
	arr1 := [2]int{1, 2} // w/o size -> declared as slice
	arr2 := arr1
	arr2[0] = 10
	fmt.Println(arr1, arr2)


	sl := arr1[0:1]
	sl[0] = 3
	fmt.Println(sl)

	printSlices(sl)	// zero value of slice => nil

	b := make([]int, 0, 5) // len(b)=0, cap(b)=5
	b = b[:cap(b)] // len(b)=5, cap(b)=5
	b = b[1:]      // len(b)=4, cap(b)=4

	b1 := make([]int, 2, 5) //len=2 cap=5 [0 0]
    printSlices(b1)

	// append
	b1 = append(b1, 100)
	printSlices(b1)
	b1 = append(b1, 200, 300, 400)
	printSlices(b1)

	// range
	for index, el := range b1 {
		fmt.Println(index, el)
	}

	// el := range b1 -> el is index not value
	// , el := range b1 -> skip index

	// slice of slices - 2d dynamic array
	twod := [][]int{
		[]int{1, 2},
		[]int{3, 4}, //add comma at the end of line as well
	}
	fmt.Printf("len=%d cap=%d %v\n", len(twod), cap(twod), twod)
	fmt.Printf("%v\n", twod[0])


// make -> value  -> allocates memory and returns
// new  -> pointer -> just creates a pointer

// slices does not - store any data/memory 
// only references to the array
}

func printSlices(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}