package main

import (
	"./array_add"
	"fmt"
)

func main() {

	arr := [] int { 1, 2, 3, 4}

	result := array_add.ArrayAdd(arr)
	fmt.Println(result)
}

