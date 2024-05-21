package main

import (
	"fmt"

)

// type contactInfo struct{
// 	email string
// 	zipCode int
// }

// type person struct {
// 	firstName string
// 	lastName  string
// 	contactInfo
// }


func main() { 



	// alex := person{"alex", "Anderson"}
	// fmt.Println(alex)
	// var alex person
	// alex.firstName="Alex"
	// alex.lastName="Anderson"
	// fmt.Println(alex)
	// fmt.Printf("%+v",alex)

	// jim:=person{
	// 	firstName: "Jim",
	// 	lastName: "Party",
	// 	contactInfo: contactInfo{
	// 		email: "jim@gmail.com",
	// 		zipCode:94000,
	// 	},
	// }
	// // jimPointer:= &jim
	// // jimPointer.updateName("Jimmy")
	// jim.updateName("Jimmy")
	// jim.print() 
	// var colors map [string]string

		// colors:=make(map[string] string)
	colors:= map[string] string{
		"red":"#ff0000",
		"green":"#4bf745",
		"white":"#ffffff",

	}
	printMap(colors)
}
  func printMap(c map[string]string){
	for color,hex:=range c{
		fmt.Println("Hex code for ",color," is ",hex)
	}
  }
	// fmt.Println(colors)
		// delete(colors,"white")
		// fmt.Println(colors)


// func (pointerToperson *person) updateName(newFirstName string){
// 	(*pointerToperson).firstName=newFirstName
// }

// func (p person) print() {
// 	fmt.Printf("%+v",p)
// }

//we can declare a  varible in go 
// var fruit string 

//var name string="ram"
//var age int =30

//alternative way for declaring and initialiation of a varible
