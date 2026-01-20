package leetcode_test

import (
	"reflect"
	"sort"
	"testing"
)

func fourSum(nums []int, target int) [][]int {
	result := [][]int{}
	n := len(nums)
	if n < 4 {
		return result
	}

	sort.Ints(nums)

	for i := 0; i < n-3; i++ {

		// skip duplicate first number
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}

		for j := i + 1; j < n-2; j++ {

			// skip duplicate second number
			if j > i+1 && nums[j] == nums[j-1] {
				continue
			}

			left, right := j+1, n-1

			for left < right {
				sum := nums[i] + nums[j] + nums[left] + nums[right]

				if sum == target {
					result = append(result, []int{
						nums[i], nums[j], nums[left], nums[right],
					})

					// skip duplicates (third number)
					for left < right && nums[left] == nums[left+1] {
						left++
					}

					// skip duplicates (fourth number)
					for left < right && nums[right] == nums[right-1] {
						right--
					}

					left++
					right--

				} else if sum < target {
					left++
				} else {
					right--
				}
			}
		}
	}

	return result
}

func TestFourSum(t *testing.T) {
	tests := []struct {
		nums     []int
		target   int
		expected [][]int
	}{
		{
			nums:     []int{1, 0, -1, 0, -2, 2},
			target:   0,
			expected: [][]int{{-2, -1, 1, 2}, {-2, 0, 0, 2}, {-1, 0, 0, 1}},
		},
		{
			nums:     []int{2, 2, 2, 2, 2},
			target:   8,
			expected: [][]int{{2, 2, 2, 2}},
		},
	}
	for _, test := range tests {
		result := fourSum(test.nums, test.target)
		if !reflect.DeepEqual(result, test.expected) {
			t.Errorf("fourSum(%v, %d) = %v; expected %v", test.nums, test.target, result, test.expected)
		}
	}
}
