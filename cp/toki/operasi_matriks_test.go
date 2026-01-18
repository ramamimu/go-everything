package toki_test

import (
	"fmt"
	"strconv"
	"testing"
)

func reflectHorizontally(matrix [][]int) [][]int {
	// implement horizontal reflection
	for i := 0; i < len(matrix)/2; i++ {
		matrix[i], matrix[len(matrix)-1-i] = matrix[len(matrix)-1-i], matrix[i]
	}
	return matrix	
}

func reflectVertically(matrix [][]int) [][]int {
	// implement vertical reflection
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[i])/2; j++ {
			matrix[i][j], matrix[i][len(matrix[i])-1-j] = matrix[i][len(matrix[i])-1-j], matrix[i][j]
		}
	}
	return matrix
}

func rotateMatrix(matrix [][]int, degrees int) [][]int {
	// implement matrix rotation by degrees (90, 180, 270)
	// use formula
	tempMatrix := matrix
	numRotations := (degrees / 90) % 4

	for r := 0; r < numRotations; r++ {
		n := len(tempMatrix)
		m := len(tempMatrix[0])
		newMatrix := make([][]int, m)
		for i := range newMatrix {
			newMatrix[i] = make([]int, n)
		}

		for i := 0; i < n; i++ {
			for j := 0; j < m; j++ {
				newMatrix[j][n-1-i] = tempMatrix[i][j]
			}
		}
		tempMatrix = newMatrix
	}

	return tempMatrix
}

func DoOperasiMatriks(n int, m int, matrix [][]int, x int, commands []string) [][]int {
	currentMatrix := matrix
	for i := range commands {
		command := commands[i]
		switch command {
		case "_":
			// reflect horizontally
			currentMatrix = reflectHorizontally(currentMatrix)
		case "|":
			// reflect vertically
			currentMatrix = reflectVertically(currentMatrix)
		default:
			rotate, _ := strconv.Atoi(command)
			// rotate matrix by 'rotate' degrees
			currentMatrix = rotateMatrix(currentMatrix, rotate)
		}
	}

	return currentMatrix
} 

func OperasiMatriks() [][]int {
	var n, m, x int
	fmt.Scan(&n, &m, &x)

	// create matrix as slice of slices
	matrix := make([][]int, n)
	for i := 0; i < n; i++ {
		matrix[i] = make([]int, m)
		for j := 0; j < m; j++ {
			fmt.Scan(&matrix[i][j])
		}
	}

	// create commands slice
	commands := make([]string, x)
	for i := 0; i < x; i++ {
		fmt.Scan(&commands[i])
	}

	return DoOperasiMatriks(n, m, matrix, x, commands)
}


func TestOperasiMatriks(t *testing.T) {
	tests := []struct {
		name     string
		n, m, x  int
		matrix   [][]int
		commands []string
		expect   [][]int
	}{
		{
			name: "horizontal reflection",
			n:    3,
			m:    3,
			x:    1,
			matrix: [][]int{
				{1, 2, 3},
				{4, 5, 6},
				{7, 8, 9},
			},
			commands: []string{"_"},
			expect: [][]int{
				{7, 8, 9},
				{4, 5, 6},
				{1, 2, 3},
			},
		},
		{
			name: "vertical reflection",
			n:    3,
			m:    3,
			x:    1,
			matrix: [][]int{
				{1, 2, 3},
				{4, 5, 6},
				{7, 8, 9},
			},
			commands: []string{"|"},
			expect: [][]int{
				{3, 2, 1},
				{6, 5, 4},
				{9, 8, 7},
			},
		},
		{
			name: "90 degree rotation",
			n:    3,
			m:    3,
			x:    1,
			matrix: [][]int{
				{1, 2, 3},
				{4, 5, 6},
				{7, 8, 9},
			},
			commands: []string{"90"},
			expect: [][]int{
				{7, 4, 1},
				{8, 5, 2},
				{9, 6, 3},
			},
		},
		{
			name: "180 degree rotation",
			n:    2,
			m:    3,
			x:    1,
			matrix: [][]int{
				{1, 2, 3},
				{4, 5, 6},
			},
			commands: []string{"180"},
			expect: [][]int{
				{6, 5, 4},
				{3, 2, 1},
			},
		},
		{
			name: "270 degree rotation",
			n:    3,
			m:    3,
			x:    1,
			matrix: [][]int{
				{1, 2, 3},
				{4, 5, 6},
				{7, 8, 9},
			},
			commands: []string{"270"},
			expect: [][]int{
				{3, 6, 9},
				{2, 5, 8},
				{1, 4, 7},
			},
		},
		{
			name: "multiple operations: horizontal then vertical",
			n:    2,
			m:    2,
			x:    2,
			matrix: [][]int{
				{1, 2},
				{3, 4},
			},
			commands: []string{"_", "|"},
			expect: [][]int{
				{4, 3},
				{2, 1},
			},
		},
		{
			name: "multiple operations: 90 rotation twice",
			n:    2,
			m:    3,
			x:    2,
			matrix: [][]int{
				{1, 2, 3},
				{4, 5, 6},
			},
			commands: []string{"90", "90"},
			expect: [][]int{
				{6, 5, 4},
				{3, 2, 1},
			},
		},
		{
			name: "complex operations: rotate, reflect horizontal, reflect vertical",
			n:    3,
			m:    2,
			x:    3,
			matrix: [][]int{
				{1, 2},
				{3, 4},
				{5, 6},
			},
			commands: []string{"90", "_", "|"},
			expect: [][]int{
				{1, 3, 5},
				{2, 4, 6},
			},
		},
		{
			name: "single element matrix",
			n:    1,
			m:    1,
			x:    1,
			matrix: [][]int{
				{5},
			},
			commands: []string{"_"},
			expect: [][]int{
				{5},
			},
		},
		{
			name: "no operations",
			n:    2,
			m:    2,
			x:    0,
			matrix: [][]int{
				{1, 2},
				{3, 4},
			},
			commands: []string{},
			expect: [][]int{
				{1, 2},
				{3, 4},
			},
		},
		{
			name: "rectangular matrix horizontal reflection",
			n:    2,
			m:    4,
			x:    1,
			matrix: [][]int{
				{1, 2, 3, 4},
				{5, 6, 7, 8},
			},
			commands: []string{"_"},
			expect: [][]int{
				{5, 6, 7, 8},
				{1, 2, 3, 4},
			},
		},
		{
			name: "rectangular matrix 90 rotation",
			n:    2,
			m:    4,
			x:    1,
			matrix: [][]int{
				{1, 2, 3, 4},
				{5, 6, 7, 8},
			},
			commands: []string{"90"},
			expect: [][]int{
				{5, 1},
				{6, 2},
				{7, 3},
				{8, 4},
			},
		},
	}
	
	for _, test := range tests {
		fmt.Printf("Running test: %s\n", test.name)
		result := DoOperasiMatriks(test.n, test.m, test.matrix, test.x, test.commands)
		match := true
		if len(result) != len(test.expect) {
			match = false
		} else {
			for i := range result {
				if len(result[i]) != len(test.expect[i]) {
					match = false
					break
				}
				for j := range result[i] {
					if result[i][j] != test.expect[i][j] {
						match = false
						break
					}
				}
			}
		}
		if !match {
			t.Errorf("FAIL: %s\nExpected: %v\nGot: %v\n", test.name, test.expect, result)
		} else {
			t.Logf("PASS: %s\n", test.name)
		}
	}
}