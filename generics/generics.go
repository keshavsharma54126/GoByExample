package main

import (
	"fmt"
	"errors"
)

//function using generics 
func splitSlice[T any](s []T)([]T,[]T){
	mid:= len(s)/2
	return s[:mid],s[mid:]
}

func chargeForLIneItem[T lineItem](newItem T,oldItems []T,balance float64)([]T,float64,error){
	if balance-newItem.GetCost()<=0.0{
		return nil,0.0,errors.New("insufficient balance")
	}
	oldItems = append(oldItems, newItem)
	balance = balance - newItem.GetCost()
	return oldItems,balance,nil
}

type lineItem interface{
	GetCost() float64 
	GetName() string
}


func main() {
	fmt.Println("generics")
}
