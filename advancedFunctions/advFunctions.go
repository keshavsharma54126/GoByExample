package main 

import (
	"fmt"
)

func add(x,y int)int{
	return x+y
}

func mul(x,y int)int{
	return x*y
}

//these are called highorder functions basically fuctnions taKING other functions as arguments
func aggregator(a,b,c int,arithmetic func(int,int)int)int{
	return arithmetic(arithmetic(a,b),c)
}
//currying (taking a function as an input and returing a function as an input)
func selfMath(mathFunc func(int,int) int) func(int)int{
	return func(x int) int{
		return mathFunc(x,x)
	}
}

func main(){
	fmt.Println(aggregator(2,3,5,add))
	fmt.Println(aggregator(2,3,5,mul))
	squareFunc:= selfMath(mul)
	doubleFunc:= selfMath(add)

	fmt.Println(squareFunc(3))
	fmt.Println(doubleFunc(4))
}