package leetcode_test

import (
	"testing"
)

func maxArea(height []int) int {
	maxArea := 0
	l, r := 0, len(height)-1
	for l < r {
		hl, hr := height[l], height[r]
		if hl > hr {
			area := hr * (r - l)
			if area > maxArea {
				maxArea = area
			}
			r--
		} else {
			area := hl * (r - l)
			if area > maxArea {
				maxArea = area
			}
			l++
		}
	}

	return maxArea
}

func maxArea_copilot(height []int) int {
	l, r := 0, len(height)-1
	maxArea := 0

	for l < r {
		h := 0
		if height[l] < height[r] {
			h = height[l]
			area := h * (r - l)
			if area > maxArea {
				maxArea = area
			}
			l++
		} else {
			h = height[r]
			area := h * (r - l)
			if area > maxArea {
				maxArea = area
			}
			r--
		}
	}

	return maxArea
}

func TestMaxArea(t *testing.T) {
	tests := []struct {
		height []int
		expect int
	}{
		{[]int{1, 8, 6, 2, 5, 4, 8, 3, 7}, 49},
		{[]int{1, 1}, 1},
		{[]int{4, 3, 2, 1, 4}, 16},
		{[]int{1, 2, 1}, 2},
	}

	for _, test := range tests {
		result := maxArea_copilot(test.height)
		if result != test.expect {
			t.Errorf("maxArea(%v) = %d; want %d", test.height, result, test.expect)
		}
	}

	for _, test := range tests {
		result := maxArea(test.height)
		if result != test.expect {
			t.Errorf("maxArea(%v) = %d; want %d", test.height, result, test.expect)
		}
	}
}
