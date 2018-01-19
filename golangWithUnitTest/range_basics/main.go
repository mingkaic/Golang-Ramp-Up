package main

import (
	"./range_basics"
	"fmt"
)

func main () {
	nums := []int{4,5,6}
	result := range_basics.RangeBasics(nums)
	fmt.Println(result)
}