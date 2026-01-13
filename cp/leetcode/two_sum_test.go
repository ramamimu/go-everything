package leetcode_test

import "testing"

func twoSum(nums []int, target int) []int {
	// key: number, value: index
	targetMap := make(map[int]int)
	for i := range nums {
		complement := target - nums[i]
		if _, ok := targetMap[complement]; ok {
			return []int{targetMap[complement], i}
		}

		targetMap[nums[i]]=i
	}
	return nil
}

func TestTwoSum(t *testing.T) {
	tests := []struct {
		nums   []int
		target int
		expect []int
	}{
		{[]int{2, 7, 11, 15}, 9, []int{0, 1}},
		{[]int{3, 2, 4}, 6, []int{1, 2}},
		{[]int{3, 3}, 6, []int{0, 1}},
		{[]int{3, 1, 12, 3, 24, 13}, 6, []int{0, 3}},
		{[]int{3, 3, 8, 12, 2, 13}, 20, []int{2, 3}},
		{[]int{-3, 4, 1, 2, 9}, 6, []int{1, 3}}, // 4 + 2
	}

	for _, test := range tests {
		result := twoSum(test.nums, test.target)
		if len(result) != len(test.expect) {
			t.Errorf("twoSum(%v, %d) = %v; want %v", test.nums, test.target, result, test.expect)
			continue
		}
		for i := range result {
			if result[i] != test.expect[i] {
				t.Errorf("twoSum(%v, %d) = %v; want %v", test.nums, test.target, result, test.expect)
				break
			}
		}
	}
}
