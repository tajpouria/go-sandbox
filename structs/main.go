package main

import "fmt"

type personInfo struct {
	email   string
	zipCode int
}

type person struct {
	firstname string
	lastname  string
	personInfo
}

func (p person) print() {
	fmt.Printf("%+v", p)
}

func (p *person) updateName(newName string) {
	(*p).firstname = newName
}

func main() {
	p := person{
		firstname: "Foo",
		lastname:  "Bar",
		personInfo: personInfo{
			email:   "foo@foo.com",
			zipCode: 123,
		},
	}

	p.updateName("NEWNAME")

	p.print()
}
