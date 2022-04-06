package main

import (
	"fmt"
)

type contactInfo struct {
	email string
	zipCode int
}

type person struct {
	firstName string
	lastName  string
	contact contactInfo
}


func main() {
	alex := person{
		"Alex", 
		"Todorov", 
		contactInfo{ 
			"alex@gmail.com", 
			1304 },
		}
	
	slice := []string{"One", "Two"}
	updateSlice(slice)
	fmt.Println(slice) // it work, referent types don't need pointer 
	// alexPointer := &alex 
	// if we use *person type as receiver we dont need the poiner

	alex.updateName("Aleks")
	alex.print()
}

func updateSlice(s []string) {
	s[0] = "Three"
}

func (pointerToPerson *person) updateName(newFirstName string) {
	(*pointerToPerson).firstName = newFirstName
}

func (p person) print() {
	fmt.Printf("%+v", p)
}