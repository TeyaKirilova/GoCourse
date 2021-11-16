package main

import "fmt"

func Sum(x, y int) int {
	sum := x
	for i := 1; i < y; i++ {
		sum = sum + x
	}
	return sum
}

func main() {
	fmt.Println("Your value is:\n", Sum(2, 3))
	t := 5
	fmt.Println(t)
}
