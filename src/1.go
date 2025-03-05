  package main

import (
	"fmt"
	"math/rand"
)

func randomCode() {
	var n int = rand.Intn(1000) // generate a random number between 0 and 999, inclusive
	fmt.Println("The generated code is", n)
}