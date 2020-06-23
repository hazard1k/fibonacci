package main

import "fmt"

// Create a cache for numbers
var cache map[int]int

func init() {
	cache = make(map[int]int)
}

func fibonacci(number int) int {

	if value, ok := cache[number]; ok {
		return value
	}

	if number <= 1 {
		return number
	}

	return fibonacci(number-2) + fibonacci(number-1)
}

func main() {
	fmt.Printf("result for number 3 is %x\n", fibonacci(3))
	fmt.Printf("result for number 8 is %x", fibonacci(8))
}
