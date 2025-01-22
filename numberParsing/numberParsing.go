package main  

import(
	"fmt"
	"strconv"
)

func main(){
	f,_ := strconv.ParseFloat("1.344",64)
	fmt.Println(f)
	i,_:= strconv.ParseInt("1234",0,64)
	fmt.Println(i)

	d,_ := strconv.ParseInt("0x1c8",0,64)
	fmt.Println(d)

	k,_:= strconv.Atoi("134")
	fmt.Println(k)

	_,e:= strconv.Atoi("wat")
	fmt.Println(e)
}

