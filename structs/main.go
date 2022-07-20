package main

import "fmt"

type person struct {
	firstName string
	lastName  string
	age       int
	contactInfo
}

type contactInfo struct {
	email string
	zip   int
}

func main() {
	var john person

	john.firstName = "John"
	john.lastName = "Anderson"
	john.age = 30
	john.contactInfo.email = "John.Anderson@mail.com"
	john.contactInfo.zip = 1234

	jim := person{
		firstName: "Jim",
		lastName:  "Smith",
		age:       32,
		contactInfo: contactInfo{
			email: "Jim.Smith@mail.com",
			zip:   1235,
		},
	}

	jim.updateName("Jimmy")
	// fmt.Printf("%v\n", &jim)
	// fmt.Printf("%v\n", jimPointer)
	//jim.print()
}

func (pointerToPerson *person) updateName(newFirstName string) {
	(*pointerToPerson).firstName = newFirstName
}

func (p person) print() {
	fmt.Printf("%+v\n", p)
}
