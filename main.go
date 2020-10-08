package main

import "fmt"

func zero(x_ptr *int) {
	*x_ptr = 0
}

func one(x_ptr *int) {
	*x_ptr = 1
}

func main() {
	x := 5
	zero(&x)
	fmt.Println(x)

	x_ptr := new(int)
	one(x_ptr)
	fmt.Println(x_ptr)
}