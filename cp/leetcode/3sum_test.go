package leetcode_test

import (
	"sort"
	"testing"
)

func threeSum(nums []int) [][]int {
	result := [][]int{}
	n := len(nums)
	if n < 3 {
		return result
	}

	sort.Ints(nums)

	for i := 0; i < n-2; i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}

		left, right := i+1, n-1
		for left < right {
			sum := nums[i] + nums[left] + nums[right]
			if sum == 0 {
				result = append(result, []int{nums[i], nums[left], nums[right]})
				if left < right && nums[left] == nums[left+1] {
					left++
				}

				if left < right && nums[right] == nums[right-1] {
					right--
				}
				left++
				right--
			} else if sum > 0 {
				right--
			} else {
				left++
			}
		}
	}

	return result
}

func TestThreeSum(t *testing.T) {
	tests := []struct {
		nums   []int
		expect [][]int
	}{
		{
			nums:   []int{-1, 0, 1, 2, -1, -4},
			expect: [][]int{{-1, -1, 2}, {-1, 0, 1}},
		},
		{
			nums:   []int{-2, 0, 1, 1, 2},
			expect: [][]int{{-2, 1, 1}, {-2, 0, 2}},
		},
		{
			nums:   []int{},
			expect: [][]int{},
		},
		{
			nums:   []int{0, 0, 0},
			expect: [][]int{{0, 0, 0}},
		},
		{
			nums:   []int{0, 1, 1},
			expect: [][]int{},
		},
	}

	for _, test := range tests {
		result := threeSum(test.nums)
		if len(result) != len(test.expect) {
			t.Errorf("threeSum(%v) = %v; want %v", test.nums, result, test.expect)
			continue
		}
	}
}
