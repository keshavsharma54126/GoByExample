package main

import (
	"fmt"
	"maps"
)


func main(){
	m := make(map[string] int)
	m["k1"]=3
	m["k2"]=2

	fmt.Println("map: ",m)
	
	v1:= m["k1"]
	fmt.Println("v1: ",v1)

	v3 := m["k2"]
	fmt.Println("v2: ",v3)
	
	fmt.Println("length of map m: ",len(m))

	delete(m,"k2")
	fmt.Println("map m: ",m)

	clear(m)
	fmt.Println("map m: ",m)

	_, prs := m["k2"]
	fmt.Println("prs: ",prs)

	n:= map[string]int{"foo":1,"bar":2}
	fmt.Println("new map: ",n)

	n2:= map[string]int{"foo":1,"bar":2}
	fmt.Println(n2)
	if maps.Equal(n,n2){
		fmt.Println("n==n2")
	}


}