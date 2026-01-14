package leetcode_test

import (
	"sort"
	"testing"
)

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	m := len(nums1) // 3
	n := len(nums2) // 4

	arr := []int{}
	arr = append(arr, nums1...)
	arr = append(arr, nums2...)

	sort.Ints(arr)

	mid := (m + n) / 2
	if (m+n)%2 == 0 {
		return float64(arr[mid-1]+arr[mid]) / 2.0
	} else {
		return float64(arr[mid])
	}
}

func TestFindMedianSortedArrays(t *testing.T) {
	tests := []struct {
		nums1  []int
		nums2  []int
		expect float64
	}{
		{[]int{1, 3}, []int{2}, 2.0},
		{[]int{1, 2}, []int{3, 4}, 2.5},
		{[]int{0, 0}, []int{0, 0}, 0.0},
		{[]int{}, []int{1}, 1.0},
		{[]int{2}, []int{}, 2.0},
	}

	for _, test := range tests {
		result := findMedianSortedArrays(test.nums1, test.nums2)
		if result != test.expect {
			t.Errorf("findMedianSortedArrays(%v, %v) = %v; want %v", test.nums1, test.nums2, result, test.expect)
		}
	}
}