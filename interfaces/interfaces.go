package main

import (
	"fmt"
	"math"
)

type shape interface{
    area() float64
    perimeter() float64
}

type Rectangle struct{
    width,height float64
}

func(r Rectangle) area()float64{
    return r.height*r.height
}
func (r Rectangle) perimeter()float64{
    return 2*r.height + 2*r.width
}

type Circle struct{
    radius float64
}

func (c Circle) area()float64{
    return math.Pi*c.radius*c.radius
}

func (c Circle) perimeter()float64{
    return 2*math.Pi*c.radius
}

func PrintShapeDetails(s shape){
    fmt.Printf("Area: ",s.area())
    fmt.Printf("perimeter: ",s.perimeter())
}

func PrintValue(val interface{}){
    fmt.Println("Value: ", val)
}
func main(){
	rectanle := Rectangle{width:2,height:2}
    circle:= Circle{radius: 2}
    PrintShapeDetails(rectanle)
    PrintShapeDetails(circle)
    PrintValue(23)
    PrintValue("hello")
    PrintValue(Circle{4.0})

}