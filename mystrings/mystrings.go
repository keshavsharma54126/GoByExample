package mystrings 

// my strings package and the function name is Reverse with capital start letter
func Reverse(s string)string{
	result:=""
	for _,v:= range s{
		result = string(v)+result
	}
	return result
}