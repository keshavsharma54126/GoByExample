package main  

import (
	"os"
	"io"
)
//defer is basically a keyword used to defer a function call at the end of a function basically whever the function will return before that the deffered function will be called 
func copyFile(dstName,srcName string)(written int64, err error){
	//open the source file
	src,err:= os.Open(srcName)
	if err!=nil{
		return
	}
	// close the srouce file when the copyfile function returns
	defer src.Close()

	//create a destination file 
	dst,err:= os.Create(dstName)
	if err!=nil{
		return 
	}

	//close the destinatiuon file when the copy file function returns
	defer dst.Close()

	return io.Copy(dst,src)


}