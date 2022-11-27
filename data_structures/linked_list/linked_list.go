package linked_list

import (
	"fmt"
	"strings"
)

type LinkedList[T any] struct {
	Head *Node[T]
	Tail *Node[T]
}

type Node[T any] struct {
	Val  T
	Next *Node[T]
	Prev *Node[T]
}

func FromSlice[T any](slice []T) *LinkedList[T] {
	head := &Node[T]{
		Val: slice[0],
	}

	prev := head

	for i := 1; i < len(slice); i++ {
		curr := &Node[T]{
			Val:  slice[i],
			Prev: prev,
		}

		prev.Next = curr

		prev = curr
	}

	return &LinkedList[T]{
		Head: head,
		Tail: prev,
	}
}

func (ll LinkedList[T]) String() string {
	var b strings.Builder

	cur := ll.Head
	for cur != nil {
		fmt.Fprintf(&b, "%+v\n", cur)
		cur = cur.Next
	}

	return b.String()
}
