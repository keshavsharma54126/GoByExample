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

func main(){
	var s shapee
	s = circlee{radius: 2.0}
	c,ok:= s.(circlee)
	if !ok {
		log.Fatal("s is not a circle")
	}

	radius := c.radius
	fmt.Println(radius)
}