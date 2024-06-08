package main

import (
	"testing"
)

func TestBinarySearch(t *testing.T) {
	tests := []struct {
		arr    []int
		target int
		want   int
	}{
		{[]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, 5, 4},
		{[]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, 1, 0},
		{[]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, 10, 9},
		{[]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, 11, -1},
		{[]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, -1, -1},
		{[]int{}, 1, -1},
	}
	for _, tt := range tests {
		if got := BinarySearch(tt.arr, tt.target); got != tt.want {
			t.Errorf("BinarySearch() = %v, want %v", got, tt.want)
		}
	}
}
