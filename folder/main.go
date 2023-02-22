package main

import "fmt"

type person struct {
	firstName string
	lastName  string
	age       int
	genre     string
}

func main() {
	lessi := person{"Lively", "Brooke", 24, "female"}
	ashy := person{
		firstName: "Ashton",
		lastName:  "Brooke",
		age:       28,
		genre:     "male",
	}
	fmt.Println(lessi, ashy)
}
