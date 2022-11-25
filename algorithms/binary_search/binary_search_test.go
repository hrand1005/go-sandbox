package main

import (
	"testing"
)

func TestBinarySearch(t *testing.T) {
	tests := []struct {
		name              string
		orderedCollection []int
		target            int
		want              int
	}{
		{
			name:              "Nominal case returns valid index",
			orderedCollection: []int{-2, -2, 0, 1, 3, 45, 69, 70},
			target:            69,
			want:              6,
		},
		{
			name:              "Not found element returns -1",
			orderedCollection: []int{-2, -2, 0, 1, 3, 45, 69, 70},
			target:            68,
			want:              NotFound,
		},
		{
			name:              "Not found element returns -1",
			orderedCollection: []int{-2, -1},
			target:            -1,
			want:              1,
		},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := BinarySearch(tc.target, tc.orderedCollection)
			if got != tc.want {
				t.Fatalf("GOT: %v\nWANT: %v\n", got, tc.want)
			}
		})
	}
}

func TestBinarySearch2(t *testing.T) {
	tests := []struct {
		name              string
		orderedCollection []int
		target            int
		want              int
	}{
		{
			name:              "Nominal case returns valid index",
			orderedCollection: []int{-2, -2, 0, 1, 3, 45, 69, 70},
			target:            69,
			want:              6,
		},
		{
			name:              "Not found element returns -1",
			orderedCollection: []int{-2, -2, 0, 1, 3, 45, 69, 70},
			target:            68,
			want:              NotFound,
		},
		{
			name:              "Not found element returns -1",
			orderedCollection: []int{-2, -1},
			target:            -1,
			want:              1,
		},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := BinarySearch2(tc.target, tc.orderedCollection)
			if got != tc.want {
				t.Fatalf("GOT: %v\nWANT: %v\n", got, tc.want)
			}
		})
	}
}
