package main 

import (
	"fmt"
	s "strings"
)

var p = fmt.Println
func main(){
	p("Contains: ", s.Contains("test","es"))
	p("Count: ",s.Count("test","t"))
	p("HasPrefix",s.HasPrefix("test","te"))
	p("hasSuffix",s.HasSuffix("test","st"))
	p("index: ",s.Index("test","e"))
	p("Join: ",s.Join([]string{"a","b"},"-"))
	p("Join: ",s.Join(([]string{"c","d"}),""))
	p("Repeat: ",s.Repeat("a",5))
	p("Replace: ",s.Replace("foo","o","0",-1))
	p("Replace: ",s.Replace("foo","o","0",1))
	p("split: ",s.Split("a-b-c-d","-"))
	p("ToLower",s.ToLower("TEST"))
	p("ToUpper",s.ToUpper("test"))
}