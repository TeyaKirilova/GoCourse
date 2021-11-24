package main

import (
	"fmt"
)

func findWinner(n, m int) int {

	mySlice := make([]int, n)
	for i := 0; i < n; i++ {
		mySlice[i] = i + 1
	}

	for index, popped := 0, 1; len(mySlice) > 1; {
		if index == len(mySlice) {
			index = 0
		}
		if m == popped {
			copy(mySlice[index:], mySlice[index+1:])
			mySlice[len(mySlice)-1] = 0
			mySlice = mySlice[:len(mySlice)-1]
			popped = 1
		} else {
			index++
			popped++
		}
	}
	p := mySlice[0]

	return p
}

func main() {
	var n int
	var m int

	fmt.Println("Enter n:")
	fmt.Scanln(&n)

	fmt.Println("Enter m:")
	fmt.Scanln(&m)

	fmt.Println("The winner is:")
	fmt.Println(findWinner(n, m))
}
