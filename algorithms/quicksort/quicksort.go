package main

type sortable interface {
	uint8 | uint16 | uint32 | uint64 |
		int | int8 | int16 | int32 | int64 |
		float32 | float64 |
		string
}

// quicksort can sort a slice of sortable elements
func quicksort[T sortable](slice []T, low, high int) {
	if low >= high {
		return
	}

	p := partition(slice, low, high)
	quicksort(slice, low, p-1)
	quicksort(slice, p+1, high)
}

// partition is responsible for ensuring some pivot index is
// put in its final position. It returns that index
func partition[T sortable](slice []T, low, high int) int {
	pivot := slice[high]
	i := low

	for j := low; j < high; j++ {
		if slice[j] < pivot {
			slice[j], slice[i] = slice[i], slice[j]
			i++
		}
	}

	slice[high], slice[i] = slice[i], slice[high]

	return i
}
