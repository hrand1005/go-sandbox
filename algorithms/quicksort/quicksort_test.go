package main

import (
	"reflect"
	"testing"
)

func TestQuickSort(t *testing.T) {
	tests := []struct {
		name  string
		input []int
		want  []int
	}{
		{
			name:  "Nominal slice",
			input: []int{1, 4, 2},
			want:  []int{1, 2, 4},
		},
		{
			name:  "Repeated elements",
			input: []int{1, 2, 1},
			want:  []int{1, 1, 2},
		},
		{
			name:  "No change",
			input: []int{1, 2, 3, 4},
			want:  []int{1, 2, 3, 4},
		},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			// sort the input in place
			quicksort(tc.input, 0, len(tc.input)-1)
			if !reflect.DeepEqual(tc.input, tc.want) {
				t.Fatalf("got: %#v\nexpected: %#v\n", tc.input, tc.want)
			}
		})
	}
}
