package others_test

import "sort"

type CustomSorter struct {
	X int
	Y int
	Z int
}

func CustomSort(arr []CustomSorter) []CustomSorter {
	sort.Slice(arr, func(i, j int) bool {
		if arr[i].X != arr[j].X {
			return arr[i].X < arr[j].X
		}
		if arr[i].Y != arr[j].Y {
			return arr[i].Y < arr[j].Y
		}
		return arr[i].Z < arr[j].Z
	})
	return arr
}
