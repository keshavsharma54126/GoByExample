package main

import(
    "fmt"

)

type point struct{
    x,y int
}

func main() {
    p:= point{1,2}
    fmt.Println(p)

    fmt.Printf("struct1: %v\n",p)
    fmt.Printf("struct2: %+v\n",p)
    fmt.Printf("struct3: %#v\n",p)
    fmt.Printf("type: %T\n",p)
    fmt.Printf("bool: %t\n",true)
    

}


