package main

func main() {
	// EXERCISE #1
	// arr := [5]int{1, 2, 3, 4, 5}
	//
	// arr[0] = 2
	// arr[1] = 3
	// arr[2] = 4
	// arr[3] = 5
	// arr[4] = 6
	//
	// for _, v := range arr {
	// 	fmt.Println(v)
	// }
	//
	// fmt.Printf("%T", arr)

	// EXERCISE #2
	// arr := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51, 52}
	//
	// for i, v := range arr {
	// 	fmt.Println(i, v)
	// }

	// fmt.Printf("%T", arr)

	// EXERCISE #3

	// fmt.Println(arr[:5])
	// fmt.Println(arr[5 : len(arr)-1])
	// fmt.Println(arr[2:7])
	// fmt.Println(arr[1:6])

	// EXERCISE #4
	// x := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}
	// x = append(x, 52)
	//
	// fmt.Println(x)
	//
	// x = append(x, 53, 54, 55)
	//
	// y := []int{56, 57, 58, 59, 60}
	//
	// x = append(x, y...)
	// fmt.Println(x)

	// EXERCISe #5
	// x := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}
	//
	// x1 := x[:3]
	// x2 := x[6:]
	//
	// x = append(x1, x2...)
	//
	// fmt.Println(x)

	// EXERCISE #6
	// states := make([]string, 0, 50)
	//
	// states = append(states, "Alabama", "etc")
	//
	// for i := 0; i < len(states); i++ {
	// 	fmt.Println(i, states[i])
	// }
	// fmt.Println(len(states), cap(states))

	// EXERCISE #7
	// x := [][]string{{"James", "Bond", "Shaken, not stirred"}, {"Miss", "Moneypenny", "Hellooooo, James."}}
	//
	// for _, v1 := range x {
	// 	for _, v2 := range v1 {
	// 		fmt.Println(v2)
	// 	}
	// }

	// EXERCISE #8
	// x := map[string][]string{
	// 	"bond_james":      []string{"Shaken not stirred", "Martinis", "Women"},
	// 	"moneypenny_miss": []string{"James Bond", "Literature", "Computer Science"},
	// 	"no_dr":           []string{"Being evil", "Ice cream", "Sunsets"},
	// }
	//
	// for k, v := range x {
	// 	fmt.Println(k)
	//
	// 	for _, thing := range v {
	// 		fmt.Println(thing)
	// 	}
	// }

	// EXERCISE #9
	// x["me_dude"] = []string{"pizza", "chicken", "sleeping"}
	//
	// for k, v := range x {
	// 	fmt.Println(k)
	//
	// 	for _, thing := range v {
	// 		fmt.Println(thing)
	// 	}
	// }

	// EXERCISE #10
	// delete(x, "me_dude")
}
