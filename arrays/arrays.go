package main

import "fmt"

func main() {
	var a [6]int
	fmt.Println("emp:", a)

	a[4]= 100
	fmt.Println("set: ",a)
	fmt.Println("get: ",a[4])

	fmt.Println("len: ",len(a))

	var name = "keshav"
	fmt.Println("len of string: ",len(name))

	b := [5] int{1,2,3,4,5}
	fmt.Println(b)

	b = [...]int{100,3:444,234}
	fmt.Println(b)

	var twoD [2][3]int
	for i:=0 ;i<len(twoD);i++{
		for j:=0;j<len(twoD[i]);j++{
			twoD[i][j]=i+j
		}
	}
	fmt.Println("2d array: ",twoD)

	twoD = [2][3] int{
		{1,2,3},
		{1,2,3},
	}
	fmt.Println("2d array reinitialized: ",twoD)

}


