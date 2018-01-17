package main

import (
	"./sortStringArray"
	"fmt"
)

func main() {
	s := []string{"c", "a", "b"}
	result := sortStringArray.SortString(s)
	fmt.Println(result)
}