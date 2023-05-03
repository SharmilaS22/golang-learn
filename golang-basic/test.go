package main

// factored import statement
import (
	"fmt"
	"math"
	"math/rand"
)

// if func - with same name use alias
// import (
// 	mathRand "math/rand"
// )

// can also be like
// import "math"
// import "math/rand"

func main1() {
	fmt.Println("Hello Sharmi!!✨")
	fmt.Println("Just the square root of 64: ", math.Sqrt(64))
	fmt.Println("Just a random number: ", rand.Intn(24))

	// All the exported names starts with uppercase letters
	// Here eg, Println, Sqrt, Intn
	// This returns error
	// fmt.Println("Hello, ", math.pi)
	// This returns Pi value
	fmt.Println("Hello, ", math.Pi)
}

// to download packages from go.mod
// go mod download