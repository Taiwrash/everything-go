package main

import "fmt"

type contactInfo struct {
	email string
	phone string
}
type person struct {
	firstname string
	lastname  string
	contactInfo
}

func main() {
	jim := person{
		firstname: "Jim",
		lastname:  "Org",
		contactInfo: contactInfo{
			email: "jim@j.co",
			phone: "+2348765445",
		},
	}
	fmt.Println(jim)
	jim.update("Jimmy")
	fmt.Println(jim)

}

func (p *person) update(firstname string) {
	(*p).firstname = firstname
}
