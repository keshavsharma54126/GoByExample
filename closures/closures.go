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

func doMath(f func(int)int,nums []int)[]int{
	var results []int 
	for _,n:=range nums{
		results = append(results,f(n) )
	}
	return results
}


func main() {
	harryPotterAggregator:= concatter()
	harryPotterAggregator("Mr.")
	harryPotterAggregator("and")
	harryPotterAggregator("Mrs.")
	harryPotterAggregator("Dursley")

	fmt.Println(harryPotterAggregator("are stupid"))
	nums:=[]int{1,2,3,4}
	//example of anonymous functions
	allNumsDoubled:= doMath(func(x int)int{
		return x*x
	},nums)
	fmt.Println(allNumsDoubled)
}
 