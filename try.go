package main

import "fmt"

type person struct {
	first string
	last  string
	age   int
}

type secretAgent struct {
	person
	ltk bool
}

func main() {
	sa1 := secretAgent{
		person: person{
			first: "James",
			last:  "Bond",
		},
		ltk: true,
	}

	p2 := person{
		first: "Miss",
		last:  "Moneypenny",
	}

	fmt.Println(sa1)
	fmt.Println(p2)
	fmt.Println(sa1.first)
}
