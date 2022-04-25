package main

type person struct {
	firstName        string
	lastName         string
	iceCreamFlavours []string
}

func main() {
	// EXERCISE #1
	// p1 := person{
	// 	firstName:        "John",
	// 	lastName:         "Doe",
	// 	iceCreamFlavours: []string{"Strawberry", "Chocolate"},
	// }
	// p2 := person{
	// 	firstName:        "Mary",
	// 	lastName:         "Ruby",
	// 	iceCreamFlavours: []string{"Vanilla", "Orange"},
	// }

	// fmt.Println(p1.firstName)
	// fmt.Println(p1.lastName)
	// for k, v := range p1.iceCreamFlavours {
	// 	fmt.Println(k, v)
	// }
	//
	// fmt.Println(p2.firstName)
	// fmt.Println(p2.lastName)
	// for k, v := range p2.iceCreamFlavours {
	// 	fmt.Println(k, v)
	// }

	// EXERCISE #2
	// personsMap := map[string]person{p1.lastName: p1, p2.lastName: p2}
	//
	// for _, v := range personsMap {
	// 	fmt.Println(v.firstName)
	// 	fmt.Println(v.lastName)
	// 	for _, v := range v.iceCreamFlavours {
	// 		fmt.Println(v)
	// 	}
	// }

	// EXERCISE #3

	// type vehicle struct {
	// 	doors int
	// 	color string
	// }
	//
	// type truck struct {
	// 	vehicle
	// 	fourWheel bool
	// }
	//
	// type sedan struct {
	// 	vehicle
	// 	luxury bool
	// }
	//
	// t := truck{
	// 	vehicle: vehicle{
	// 		doors: 5,
	// 		color: "Red",
	// 	},
	// 	fourWheel: true,
	// }
	//
	// s := sedan{
	// 	vehicle: vehicle{
	// 		doors: 10,
	// 		color: "Black",
	// 	},
	// 	luxury: false,
	// }
	//
	// fmt.Println(t.doors)
	// fmt.Println(t.color)
	// fmt.Println(t.fourWheel)
	//
	// fmt.Println(s.doors)
	// fmt.Println(s.color)
	// fmt.Println(s.luxury)

	// EXERCISE #4
	// v := struct {
	// 	t       string
	// 	options []int
	// }{
	// 	t: "Foobar",
	// 	options: []int{
	// 		5,
	// 		10,
	// 		12,
	// 	},
	// }
	//
	// fmt.Println(v)

}
