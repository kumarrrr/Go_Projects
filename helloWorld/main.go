package main

import "fmt"

func main() {
	fmt.Println("Initial level of understanding given")
	//variable defining
	var name string
	name = "go training"
	fmt.Println("Variable declaration understood", name)

	// calling external function inside & assigning it to a new variable,(age), in a different way
	age := myAge()

	fmt.Println(age)

}

// defining another function
func myAge() string {
	return "age is just a number!Keep Learning"
}
