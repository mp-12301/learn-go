package main

import "fmt"

func main() {
	// m := map[string]int{
	// 	"James":           32,
	// 	"Miss Moneypenny": 27,
	// }
	// fmt.Println(m)
	// fmt.Println(m["James"])
	// fmt.Println(m["Barnabas"])
	//
	// v, ok := m["Bananas"]
	// fmt.Println(v)
	// fmt.Println(ok)
	//
	// m["todd"] = 33
	//
	// if v, ok := m["James"]; ok {
	// 	fmt.Println("THIS IS THE PRINT", v)
	// }
	//
	// for k, v := range m {
	// 	fmt.Println(k, v)
	// }
	//
	// xi := []int{1, 2, 4, 6, 8}
	//
	// for i, v := range xi {
	// 	fmt.Println(i, v)
	// }
	//
	// delete(m, "James")
	// fmt.Println(m)
	//
	// delete(m, "Derp")
	// fmt.Println(m)
	//
	// if v, ok := m["Miss Moneypenny"]; ok {
	// 	fmt.Println("value:", v)
	// 	delete(m, "Miss Moneypenny")
	// }
	//
	// fmt.Println(m)

	arr := [5]int{1, 2, 3, 4, 5}

	s := arr[:]
	fmt.Println(s, len(s), cap(s))

	s1 := append(s, 99)
	fmt.Println(s1, len(s1), cap(s1))

	s1[0] = 19
	fmt.Println(s1, len(s1), cap(s1))

	fmt.Println(s, len(s), cap(s))
	// s = s[:3]
	// fmt.Println(s, len(s), cap(s))

	// s = s[0:4]
	// fmt.Println(s, len(s), cap(s))
}
