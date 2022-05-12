package main

import (
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"
)

// EXERCISE #2
// type person struct {
// 	first string
// 	last  string
// }
//
// type human interface {
// 	speak() string
// }
//
// func (p *person) speak() string {
// 	say := "my name is " + p.first
//
// 	p.first = "Foo already spoke"
//
// 	return say
// }
//
// func saySomething(h human) {
// 	fmt.Println(h.speak())
// }
//
// type T struct {
// 	i int
// }
//
// type ValueMethodCaller interface {
// 	valueMethod()
// }
//
// type PointMethodCaller interface {
// 	pointerMethod()
// }
//
// func callValueMethodOnInterface(v ValueMethodCaller) {
// 	v.valueMethod()
// }
//
// func callPointerMethodINterface(p PointMethodCaller) {
// 	p.pointerMethod()
// }
//
// // Pointer type receiver
// func (receiver *T) pointerMethod() {
// 	fmt.Printf("Pointer method called on \t%#v with address %p\n\n", *receiver, receiver)
// }
//
// // Value type receiver
// func (receiver T) valueMethod() {
// 	fmt.Printf("Value method called on \t%#v with address %p\n\n", receiver, &receiver)
// }

type SafeCounter struct {
	mu      sync.Mutex
	counter int
}

func main() {

	// EXERCISE #1
	// var wg sync.WaitGroup
	// wg.Add(2)
	//
	// func() {
	// 	fmt.Println("Im the main goroutine")
	// }()
	//
	// go func() {
	// 	fmt.Println("Im the first goroutine")
	// 	wg.Done()
	// }()
	//
	// go func() {
	// 	fmt.Println("Im the second goroutine")
	// 	wg.Done()
	// }()

	// wg.Wait()

	// EXERCISE #2
	// p := person{
	// 	first: "Foo",
	// 	last:  "Bar",
	// }
	//
	// saySomething(&p)
	// saySomething(&p)
	// var (
	// 	val     T  = T{}
	// 	pointer *T = &val
	// )
	// fmt.Printf("Value created \t%#v with address %p\n", val, &val)
	// fmt.Printf("Pointer created on \t%#v with address %p\n", *pointer, pointer)
	// fmt.Println()
	//
	// val.valueMethod()
	// pointer.pointerMethod()
	//
	// callValueMethodOnInterface(val)
	// callPointerMethodINterface(pointer)

	// EXERCISE #3
	// var wg sync.WaitGroup
	// wg.Add(100)
	//
	// counter := 0
	// newCounter := 0
	//
	// for i := 0; i < 100; i++ {
	// go func() {
	// 	newCounter = counter
	// 	runtime.Gosched()
	//
	// 	newCounter++
	// 	counter = newCounter
	// 	wg.Done()
	// }()
	// }
	//
	// wg.Wait()
	// fmt.Println(counter)

	// EXERCISE #4
	// var wg sync.WaitGroup
	// wg.Add(100)
	//
	// c := SafeCounter{counter: 0}
	// newCounter := 0
	//
	// for i := 0; i < 100; i++ {
	// 	go func() {
	// 		c.mu.Lock()
	// 		newCounter = c.counter
	// 		runtime.Gosched()
	//
	// 		newCounter++
	// 		c.counter = newCounter
	// 		c.mu.Unlock()
	// 		wg.Done()
	// 	}()
	// }
	//
	// wg.Wait()
	// fmt.Println(c.counter)

	// EXERCISE #5
	var wg sync.WaitGroup
	wg.Add(100)

	var counter uint64

	for i := 0; i < 100; i++ {
		go func() {
			runtime.Gosched()
			atomic.AddUint64(&counter, 1)
			wg.Done()
		}()
	}

	wg.Wait()
	fmt.Println(counter)
}
