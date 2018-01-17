package sortstringarray

import "sort"

func SortString(s []string) []string  {
	sort.Strings(s)
	return s
}