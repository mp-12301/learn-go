package main

import "fmt"

// EXERCISE #2
type person struct {
	first string
	last  string
}

func changeMe(p *person) {
	(*p).first += " changed"
	(*p).last += " changed"
}

func main() {
	// EXERCISE #1
	// foo := "bar"
	// fmt.Println(&foo)

	// EXERCISE #2
	p := person{
		first: "John",
		last:  "Doe",
	}
	fmt.Println(p)

	changeMe(&p)

	fmt.Println(p)
}
