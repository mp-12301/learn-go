package main

// EXERCISE #1
// func foo() int {
// 	return 5
// }
//
// func bar() (int, string) {
// 	return 5, "foobar"
// }

// EXERCISE #2
// func sum(a []int) int {
// 	sum := 0
// 	for _, v := range a {
// 		sum += v
// 	}
// 	return sum
// }
//
// func foo(x ...int) int {
// 	return sum(x)
// }
//
// func bar(p []int) int {
// 	return sum(p)
// }

// EXERCISE #3
// func foo() {
// 	fmt.Println("executed at the end")
// }
// func bar() {
// 	fmt.Println("I am the first to run")
// }

// EXERCISE #4
// type person struct {
// 	first string
// 	last  string
// 	age   int
// }
//
// func (p person) speak() {
// 	fmt.Println("My name is ", p.first, p.last, " and my age is ", p.age)
// }

// EXERCISE #5
// type square struct {
// 	width, height float64
// }
//
// type circle struct {
// 	radius float64
// }
//
// func (s square) area() float64 {
// 	return s.width * s.height
// }
//
// func (c circle) area() float64 {
// 	return c.radius * math.Pi * 2
// }
//
// type shape interface {
// 	area() float64
// }
//
// func info(s shape) {
// 	fmt.Println("Area ", s.area())
// }

// EXERCISE #8
// func createGenericFn() func() int {
// 	return func() int {
// 		return 42
// 	}
// }

func main() {
	// EXERCISE #1
	// fmt.Println(foo())
	// fmt.Println(bar())

	// EXERCISE #2
	// s := []int{1, 2, 3, 4, 5}
	// fmt.Println(foo(s...))
	// fmt.Println(bar(s))

	// EXERCISE #3
	// defer foo()
	// bar()

	// EXERCISE #4
	// p := person{
	// 	first: "John",
	// 	last:  "Doe",
	// 	age:   21,
	// }
	// p.speak()

	// EXERCISE #5
	// s := square{
	// 	width:  15,
	// 	height: 15,
	// }
	// c := circle{
	// 	radius: 20,
	// }
	//
	// info(s)
	// info(c)
	// fmt.Printf("%T", s.area)

	// EXERCISE #6
	// r := func() int {
	// 	return 42
	// }()
	//
	// fmt.Println(r)

	// EXERCISE #7
	// f := func() int {
	// 	return 42
	// }
	// fmt.Println(f())

	// EXERCISE #8
	// f := createGenericFn()
	// fmt.Println(f())
	// fmt.Printf("%T", f)

	// EXERCISE #9
	// cb := func() int {
	// 	return 42
	// }
	// fn := func(f func() int) int {
	// 	return f() - 42
	// }
	//
	// fmt.Println(fn(cb))

}
