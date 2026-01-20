package leetcode_test

import (
	"sort"
	"testing"
)

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func threeSumClosest(nums []int, target int) int {
	n := len(nums)

	// Handle edge case: if we have exactly 3 or fewer numbers, return their sum
	if n < 3 {
		val := 0
		for i := range nums {
			val += nums[i]
		}
		return val
	}

	sort.Ints(nums)
	closest := nums[0] + nums[1] + nums[2] // Initialize with first three elements

	for i := 0; i < n-2; i++ {
		// Skip duplicate values for optimization
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}

		l, r := i+1, n-1
		for l < r {
			sum := nums[i] + nums[l] + nums[r]

			// If exact match found, return immediately
			if sum == target {
				return sum
			}

			// Update closest if current sum is closer to target
			if abs(sum-target) < abs(closest-target) {
				closest = sum
			}

			// Move pointers based on comparison with target
			if sum < target {
				l++
			} else {
				r--
			}
		}
	}

	return closest
}

func threeSumClosest_copilot(nums []int, target int) int {
	n := len(nums)
	if n < 3 {
		sum := 0
		for _, num := range nums {
			sum += num
		}
		return sum
	}

	sort.Ints(nums)
	closest := nums[0] + nums[1] + nums[2]

	for i := 0; i < n-2; i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}

		left, right := i+1, n-1
		for left < right {
			sum := nums[i] + nums[left] + nums[right]
			if sum == target {
				return sum
			}

			if abs(sum-target) < abs(closest-target) {
				closest = sum
			}

			if sum < target {
				left++
			} else {
				right--
			}
		}
	}

	return closest
}

func TestThreeSumClosest(t *testing.T) {
	tests := []struct {
		name   string
		nums   []int
		target int
		expect int
	}{
		{
			name:   "basic case from leetcode",
			nums:   []int{-1, 2, 1, -4},
			target: 1,
			expect: 2,
		},
		{
			name:   "all zeros",
			nums:   []int{0, 0, 0},
			target: 1,
			expect: 0,
		},
		{
			name:   "exact match exists",
			nums:   []int{1, 1, 1, 0},
			target: 3,
			expect: 3,
		},
		{
			name:   "positive numbers",
			nums:   []int{3, 9, 1, 0},
			target: 3,
			expect: 4,
		},
		{
			name:   "negative numbers",
			nums:   []int{-5, -3, -1, 0, 2},
			target: -2,
			expect: -2, // -3 + -1 + 2 = -2 (exact match)
		},
		{
			name:   "exactly 3 elements",
			nums:   []int{1, 2, 3},
			target: 10,
			expect: 6,
		},
		{
			name:   "large positive target",
			nums:   []int{1, 2, 4, 8, 16, 32, 64, 128},
			target: 82,
			expect: 82,
		},
		{
			name:   "duplicates in array",
			nums:   []int{1, 1, 1, 1},
			target: 0,
			expect: 3,
		},
		{
			name:   "mixed positive and negative",
			nums:   []int{-100, -98, -2, -1, 1, 2, 3, 4},
			target: 0,
			expect: 0,
		},
		{
			name:   "larger array",
			nums:   []int{-3, -2, -5, 3, -4, 1, 2, 5},
			target: 1,
			expect: 1,
		},
		{
			name:   "target far from possible sums",
			nums:   []int{0, 1, 2},
			target: 100,
			expect: 3,
		},
		{
			name:   "negative target",
			nums:   []int{-10, -5, -2, 0, 3, 7},
			target: -15,
			expect: -15, // -10 + -5 + 0 = -15 (exact match)
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := threeSumClosest(test.nums, test.target)
			if result != test.expect {
				t.Errorf("FAIL: %s\nExpected: %d\nGot: %d\nTarget: %d\nNums: %v",
					test.name, test.expect, result, test.target, test.nums)
			} else {
				t.Logf("PASS: %s (result=%d)", test.name, result)
			}
		})

		t.Run(test.name+" Copilot", func(t *testing.T) {
			result := threeSumClosest_copilot(test.nums, test.target)
			if result != test.expect {
				t.Errorf("FAIL Copilot: %s\nExpected: %d\nGot: %d\nTarget: %d\nNums: %v",
					test.name, test.expect, result, test.target, test.nums)
			} else {
				t.Logf("PASS Copilot: %s (result=%d)", test.name, result)
			}
		})
	}
}
