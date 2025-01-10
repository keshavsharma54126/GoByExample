package main 

import(
	"fmt"
)

func plus(a int,b int)int{
	return a+b
}
func plusPlus(a int,b int,c int)int {
	return a+b+c
}

func addArray(a[10]int)int{
	var sum int
	for i:=0;i<len(a);i++{
		sum+=a[i]
	}
	return sum
}

func main(){
	res:= plus(1,2)
	fmt.Println(res)

	res = plusPlus(1,2,3)
	fmt.Println(res)
	
	array := [10]int{1,2,5,23,23,324,234,234,523,23}
	res2 := addArray(array)
	fmt.Println(res2)
}