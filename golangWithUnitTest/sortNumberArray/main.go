package main

import (
	"./sortNumberArray"
	"fmt"
)

func main () {
	a := [] int{5,3,4,1,2}
	result := sortNumberArray.SortArray(a)
	fmt.Println(result)
}
