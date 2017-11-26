package main

import (
	"fmt"
)

//Execute function to get price of the call option
func main() {
	r := OptInput{100.0, 105.0, 10.0, 0.10, 0.20, 1000000}
	fmt.Println(r.call())
}
