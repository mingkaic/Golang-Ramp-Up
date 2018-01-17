package sortNumberArray

import "fmt"
import "sort"

func SortArray(a []int) []int {
	sort.Ints(a)
	fmt.Println("Ints: ", a)

	s := sort.IntsAreSorted(a)
	fmt.Println("Sorted: ", s)

	return a;
}