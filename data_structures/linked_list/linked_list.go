package linked_list

import (
	"fmt"
	"strings"
)

type Node[T any] struct {
	Val  T
	Next *Node[T]
}

func FromSlice[T any](slice []T) *Node[T] {
	dummy := &Node[T]{}
	cur := dummy

	for i := 0; i < len(slice); i++ {
		cur.Next = &Node[T]{
			Val: slice[i],
		}
		cur = cur.Next
	}

	return dummy.Next
}

func (head *Node[T]) ToSlice() []T {
	res := []T{}

	cur := head
	for cur != nil {
		res = append(res, cur.Val)
	}

	return res
}

func (head *Node[T]) ListString() string {
	var b strings.Builder

	cur := head
	for cur != nil {
		fmt.Fprintf(&b, "%+v\n", cur)
		cur = cur.Next
	}

	return b.String()
}
