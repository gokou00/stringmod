package main

import(
	"fmt"
	"github.com/gokou00/stringmod/quotes"
	"github.com/gokou00/stringmod/strings"
)

func main(){
	o , e := strings.CountOddEven("12345")
	fmt.Println(o,e)

	fmt.Println(quotes.GetEmoji("turtle"))
}

