package main

import "fmt"

func main() {
	numbers := [11]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	for _, v := range numbers {
		result := ""
		if v%2 == 0 {
			result = "even"
		} else {
			result = "odd"
		}
		fmt.Printf(" %v is: %v\n", v, result)
	}
}

/*
I mainly feel that I reinforced my basic knowledge of go and also
forced myself to think a bit further when designing tests above all.
I learned go in a bootcamp, but I want to go even further to perform
well in my job.
*/
