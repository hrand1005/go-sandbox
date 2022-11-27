package linked_list

import (
	"testing"
)

func TestFromSlice(t *testing.T) {
	tests := []struct {
		name  string
		input []int
	}{
		{
			name:  "Nominal int slice",
			input: []int{0, 1, 2, 3, 4, 5},
		},
		{
			name:  "Single element slice",
			input: []int{0},
		},
		{
			name:  "Empty slice",
			input: []int{0},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {

			linkedList := FromSlice(tc.input)

			cur := linkedList.Head
			for i, v := range tc.input {
				if cur.Val != v {
					t.Fatalf("At index %v\nGOT: %v\nWANT: %v\n\nLinked List:\n%s", i, cur.Val, v, linkedList)
				}

				cur = cur.Next
			}
		})
	}
}
