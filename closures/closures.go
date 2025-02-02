package main

import(
	"fmt"
)
// a closure is a function that references variables from outside its own function body.
// the function may access and assign to the referenced variables.

func concatter()func (string )string{
	doc:=""
	return func(word string)string{
		doc+=word+ " "
		return doc
	}
}

func main() {
	harryPotterAggregator:= concatter()
	harryPotterAggregator("Mr.")
	harryPotterAggregator("and")
	harryPotterAggregator("Mrs.")
	harryPotterAggregator("Dursley")

	fmt.Println(harryPotterAggregator("are stupid"))
}
 