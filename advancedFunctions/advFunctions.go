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

func main(){
	fmt.Println(aggregator(2,3,5,add))
	fmt.Println(aggregator(2,3,5,mul))
}