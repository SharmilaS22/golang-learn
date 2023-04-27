package main

import "fmt"
import "strings"

func main5()  {
	// like python dict
	// zero value is nil
	m := make(map[string]int)
	print(m)
	m["one"] = 1
	m["two"] = 2
	print(m)
	print(m["one"])

	type Book struct {
		title string
		author string
	}

	book1 := Book{"AGGGTM", "Holly Jackson"}
	print(book1)

	m3 := map[int]string{
		3: "tues",
		4: "wed",
	}
	print(m3)

	m1 := make(map[int]Book)
	m1[1] = book1
	print(m1)
	print(m1[1])

	m4 := map[int]Book{
		1001: {"AGGGTM", "Holly Jackson"}, //dont have to
		1002: {"GGBB",   "Holly Jackson"}, //initialise again
		1003: {"AGAD",   "Holly Jackson"}, // with Book{}
	} 
	print(m4)

	// deleting from  map
	// m  is  map[one:1 two:2]
	delete(m, "one")
	print(m)

	// is present?
	elem, isPresent := m["one"]
	fmt.Println(elem, isPresent) //isPresent -> true or false
	//elem = (zero value of that type) when not present
	elem1, isPresent1 := m["two"]
	fmt.Println(elem1, isPresent1)

	print(WordCount("hello world"))
	
}

func WordCount(s string) map[string]int {
	var words []string = strings.Split(s, " ")
	print(words)
	wordCount := make(map[string]int)
	for _, elem := range words {
		count, isPresent := wordCount[elem]
		if isPresent {
			wordCount[elem] = count + 1
		} else {
			wordCount[elem] = 1
		}
	}
	return wordCount
}

func print(val interface{}) {
	fmt.Println(val)
}