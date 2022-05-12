package main

import (
	"fmt"
	"sort"
)

// EXERCISE #1
// type user struct {
// 	First string
// 	Age   int
// }

// EXERCISE #2
// type Person struct {
// 	First   string
// 	Last    string
// 	Age     int
// 	Sayings []string `json:"Sayings"`
// }

// EXERCISE #3
// type user struct {
// 	First string
// 	Age   int
// }

// EXERCISE #5
type user struct {
	First   string
	Last    string
	Age     int
	Sayings []string
}

type ByName []user

func (users ByName) Len() int           { return len(users) }
func (users ByName) Swap(i, j int)      { users[i], users[j] = users[j], users[i] }
func (users ByName) Less(i, j int) bool { return users[i].First < users[j].First }

type ByAge []user

func (users ByAge) Len() int           { return len(users) }
func (users ByAge) Swap(i, j int)      { users[i], users[j] = users[j], users[i] }
func (users ByAge) Less(i, j int) bool { return users[i].Age < users[j].Age }

func main() {
	// EXERCISE #1
	// u1 := user{
	// 	First: "James",
	// 	Age:   32,
	// }
	//
	// u2 := user{
	// 	First: "Moneypenny",
	// 	Age:   27,
	// }
	//
	// u3 := user{
	// 	First: "M",
	// 	Age:   54,
	// }
	//
	// users := []user{u1, u2, u3}
	//
	// fmt.Println(users)
	//
	// for _, user := range users {
	// 	userJson, err := json.Marshal(user)
	// 	if err != nil {
	// 		fmt.Println("something went wrong converting user to json, this one", user)
	// 	}
	// 	fmt.Println(string(userJson))
	// }

	// EXERCISE #2

	// s := `[
	// 	{"First":"James","Last":"Bond","Age":32,"Sayings":["Shaken, not stirred","Youth is no guarantee of innovation","In his majesty's royal service"]},
	// 	{"First":"Miss","Last":"Moneypenny","Age":27,"Sayings":["James, it is soo good to see you","Would you like me to take care of that for you, James?","I would really prefer to be a secret agent myself."]},
	// 	{"First":"M","Last":"Hmmmm","Age":54,"Sayings":["Oh, James. You didn't.","Dear God, what has James done now?","Can someone please tell me where James Bond is?"]}
	// ]`
	// fmt.Println(s)
	//
	// var persons []Person
	//
	// err := json.Unmarshal([]byte(s), &persons)
	//
	// if err != nil {
	// 	fmt.Println(err)
	// }
	//
	// fmt.Println(persons)

	// EXERCISE #3
	// u1 := user{
	// 	First: "John",
	// 	Age:   32,
	// }
	// encoder := json.NewEncoder(os.Stdout)
	// encoder.Encode(u1)

	// EXERCISE #4
	// xi := []int{5, 8, 2, 43, 17, 987, 14, 12, 21, 1, 4, 2, 3, 93, 13}
	// xs := []string{"random", "rainbow", "delights", "in", "torpedo", "summers", "under", "gallantry", "fragmented", "moons", "across", "magenta"}
	//
	// fmt.Println(xi)
	// // sort xi
	// sort.Ints(xi)
	// fmt.Println(xi)
	//
	// fmt.Println(xs)
	// // sort xs
	// sort.Strings(xs)
	// fmt.Println(xs)

	// EXERCISE #5
	u1 := user{
		First: "James",
		Last:  "Bond",
		Age:   32,
		Sayings: []string{
			"Shaken, not stirred",
			"Youth is no guarantee of innovation",
			"In his majesty's royal service",
		},
	}

	u2 := user{
		First: "Miss",
		Last:  "Moneypenny",
		Age:   27,
		Sayings: []string{
			"James, it is soo good to see you",
			"Would you like me to take care of that for you, James?",
			"I would really prefer to be a secret agent myself.",
		},
	}

	u3 := user{
		First: "M",
		Last:  "Hmmmm",
		Age:   54,
		Sayings: []string{
			"Oh, James. You didn't.",
			"Dear God, what has James done now?",
			"Can someone please tell me where James Bond is?",
		},
	}

	users := []user{u1, u2, u3}
	printNicely("--------------- without changes", users)

	for _, v := range users {
		sort.Strings(v.Sayings)
	}

	sort.Sort(ByAge(users))
	printNicely("--------------- sorted by age", users)

	sort.Sort(ByName(users))
	printNicely("--------------- sorted by name", users)

}

func printNicely(label string, us []user) {
	fmt.Println(label)

	for i, v := range us {
		fmt.Println("User #", i)
		fmt.Println("\t", v.First, v.Last, v.Age)
		fmt.Println("\t", "Sayings")

		for _, v := range v.Sayings {
			fmt.Println("\t\t", v)
		}
	}

}
