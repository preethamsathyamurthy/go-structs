package main

import "fmt"

type contactInfo struct {
	email   string
	pincode int64
}

type person struct {
	firstName string
	lastName  string
	contact   contactInfo
}

func (p person) print() {
	fmt.Printf("%+v", p)
}

// * person refers that it can accept only a pointer to person
func (pointerToPerson *person) updateFirstName(fname string) {
	(*pointerToPerson).firstName = fname
	// Here * is an operator which gives the value of the pointer
}

func main() {
	preetham := person{
		firstName: "Preetham",
		lastName:  "Sathyamurthy",
		contact: contactInfo{
			email:   "preethamsathyamurthy@gmail.com",
			pincode: 638006,
		},
	}
	preetham.print()

	preethamPointer := &preetham // & is an operator which returns the address of the variable
	preethamPointer.updateFirstName("Preethu")
	preetham.print()
}
