package main

import (
	"fmt"
	"sync"
)

func someroutine (content string, wg *sync.WaitGroup) {
	fmt.Println(content)
	defer wg.Done()
}

func main()  {
	waitGroup := sync.WaitGroup{}
	waitGroup.Add(2) //no of go routines for which we are wait
	// should know no of goroutines
	go someroutine("Hello", &waitGroup)
	go someroutine("World", &waitGroup)

	waitGroup.Add(1)
	go someroutine("Meow", &waitGroup)

	waitGroup.Wait()

	fmt.Println("Hello")

	mutex := sync.Mutex{} //mutual exclusion

	stream := make(chan int)
	// bufferedStream := make(chan int, 10)

	go add(1, 2, stream)
	res := <- stream // run concurrently and wait for this to complete
	fmt.Println("Res: ", res)
	// use different stream for different usecases
	

	// multiple consumers use same channel
}


func add(val1, val2 int, stream chan int) {
	stream <- val1 + val2
}