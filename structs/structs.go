package main

import "fmt"

type person struct {
	name string
	age  int
}

type messageToSend struct {
	phoneNumber int
	message     string
}

func testMessage(m messageToSend) {
	fmt.Printf("Sending message: '%s' to: %v\n", m.message, m.phoneNumber)
}

type car struct {
	Make   string
	Model  string
	Height string
	Width  string
}

//embedded struct

type truck struct{
	car 
	bedSize int
}

func newCar(Make string, Model string, Heigh string, width string) *car {
	c := car{Make: Make, Model: Model, Height: Heigh, Width: width}
	return &c
}

func printCar(car *car) {
	fmt.Println(car.Make)
}

func newPerson(name string) *person {
	p := person{name: name}
	p.age = 5
	return &p
}
//methods on structs
type rect struct{
	width int 
	height int
}

func (r rect)area() int{
	return r.width*r.height
}

type authenticationMethod struct{
	username  string 
	password string
}

func(authI authenticationMethod)getBasicAuth()string{
	return fmt.Sprintf(
		"Authorization:Basic %s:%s",
		authI.username,
		authI.password,
	)
}
func main() {

	lanesTruck:=truck{
		bedSize: 12,
		car: car{
			Make: "asdoi",
			Model: "sadfihwer",
			Width: "123 fit",
			Height: "23ft",
		},

	}
	fmt.Println(lanesTruck)

	car1 := newCar("suzuki", "2001", "10ft", "12ft")
	printCar(car1)
	testMessage(messageToSend{
		phoneNumber: 213453453245,
		message:     "Thanks for helping me",
	})
	message := messageToSend{
		message:     "hello",
		phoneNumber: 314523452345,
	}
	testMessage(message)
	fmt.Println(person{"keshav", 23})
	fmt.Println(person{name: "Alice", age: 30})
	fmt.Println(person{name: "fred"})
	fmt.Println(&person{name: "Ann", age: 40})
	fmt.Println(newPerson("John"))
	s := person{name: "sean", age: 89}
	fmt.Println(s.name)

	sp := &s

	fmt.Println(sp.age)

	sp.age = 778
	fmt.Println(sp.age)

	//this is an anonymous struct
	dog := struct {
		name   string
		isGood bool
	}{
		"Rex",
		true,
	}

	fmt.Println(dog)
	r:=rect{
		width:2,
		height: 12,
	}
	fmt.Println(r.area())
}
