package main

import (
	"fmt"
	"log"
	"math"
)

type shapee interface{
	area() float64
}

type circlee struct{
	radius float64
}

func (c circlee) area()float64{
	return math.Pi*c.radius*c.radius
}

//another better method for type assertion using switch statmemnt
func printNumericValue(num interface{}){
	switch v:= num.(type){
	case int:
		fmt.Printf("%T\n",v)
	case string:
		fmt.Printf("%T\n",v)
	default:
		fmt.Printf("%T\n",v)
	}

}
func main(){
	var s shapee
	s = circlee{radius: 2.0}
	c,ok:= s.(circlee)
	if !ok {
		log.Fatal("s is not a circle")
	}

	radius := c.radius
	fmt.Println(radius)
	printNumericValue(1)
	printNumericValue("keshav")
	printNumericValue(struct{}{})


}