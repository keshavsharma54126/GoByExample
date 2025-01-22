package main

import "fmt"

type person struct {
	name string
	age  int
}

func newPerson(name string) *person {
	p := person{name: name}
	p.age = 5
	return &p
}

func main() {
	fmt.Println(person{"keshav", 23})
	fmt.Println(person{name: "Alice", age: 30})
	fmt.Println(person{name: "fred"})
	fmt.Println(&person{name:"Ann",age:40})
	fmt.Println(newPerson("John"))
	s:=person{name:"sean",age:89}
	fmt.Println(s.name)

	sp:=&s 

	fmt.Println(sp.age)

	sp.age = 778
	fmt.Println(sp.age)

	dog:= struct{
		name string
		isGood bool
	}{
		"Rex",
		true,
	}

	fmt.Println(dog)

}


