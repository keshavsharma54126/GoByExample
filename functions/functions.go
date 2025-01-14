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

//returning multiple values

func vals()(int,int){
	return 3,7
}

func vals2()(int,string){
	return 5,"hello"
}

func main(){
	res:= plus(1,2)
	fmt.Println(res)

	res = plusPlus(1,2,3)
	fmt.Println(res)
	
	array := [10]int{1,2,5,23,23,324,234,234,523,23}
	res2 := addArray(array)
	fmt.Println(res2)

	a,b := vals()
	fmt.Println(a,b)

	_,c := vals()
	fmt.Println(c)
	val1,val3:= vals2();
	fmt.Println(val1,val3)
}