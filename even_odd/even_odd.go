package main

import "fmt"

func makeRange(min, max int) []int {
	a := make([]int, max-min+1)
	for i := range a {
		a[i] = min + i
	}
	return a
}

func isEven(n int) bool {
	if n%2 == 0 {
		return true
	}
	return false
}

func mainAll() {
	s := makeRange(0, 10)
	for _, number := range s {
		if isEven(number) {
			fmt.Println(number, "is even")
		} else {
			fmt.Println(number, "is odd")
		}
	}
}
