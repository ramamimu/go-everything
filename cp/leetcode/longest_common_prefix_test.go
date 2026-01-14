package leetcode_test

import "testing"

func longestCommonPrefix(strs []string) string {
    commonPrefix := ""
	for i := range strs{
		if i == 0 {
			commonPrefix = strs[i]
			continue 
		}

		// pop characters from commonPrefix until it matches the start of strs[i]
		currentStr := strs[i]
		if len(commonPrefix) > len(currentStr) {
			commonPrefix = commonPrefix[:len(currentStr)]
		} else {
			currentStr = currentStr[:len(commonPrefix)]
		} 

		for  j := range currentStr {
			if commonPrefix[j] != currentStr[j] {
				commonPrefix = commonPrefix[:j]
				break
			}
		}

		if commonPrefix == "" {
			return ""
		}
	}
	return commonPrefix
}

func longestCommonPrefix_v2(strs []string) string {
		if len(strs) == 0 {
		return ""
	}

	prefix := strs[0]
	for i := 1; i < len(strs); i++ {
		j := 0
		for j < len(prefix) && j < len(strs[i]) && prefix[j] == strs[i][j] {
			j++
		}
		prefix = prefix[:j]
		if prefix == "" {
			return ""
		}
	}
	return prefix
}

func TestLongestCommonPrefix(t *testing.T) {
	type testCase struct {
		strs   []string
		expect string
	}

	tests := []testCase{
		{[]string{"flower", "flow", "flight"}, "fl"},
		{[]string{"dog", "racecar", "car"}, ""},
		{[]string{"interspecies", "interstellar", "interstate"}, "inters"},
		{[]string{"throne", "dungeon"}, ""},
		{[]string{"throne", "throne"}, "throne"},
	}

	for _, test := range tests {
		result := longestCommonPrefix(test.strs)
		if result != test.expect {
			t.Errorf("longestCommonPrefix(%v) = %q; want %q", test.strs, result, test.expect)
		}
	}	

	for _, test := range tests {
		result := longestCommonPrefix_v2(test.strs)
		if result != test.expect {
			t.Errorf("longestCommonPrefix_v2(%v) = %q; want %q", test.strs, result, test.expect)
		}
	}
}