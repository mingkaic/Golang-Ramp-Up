package array_add
import "fmt"


func ArrayAdd(arr []int) []int{

	if len(arr) == 0 {
		return nil
	}

	fmt.Println("ArrayAdd: In")
	for i := 0; i < len(arr); i++ {

		arr[i] = arr[i] + 1
	}


	fmt.Println("ArrayAdd: Out", arr)
	return arr
}