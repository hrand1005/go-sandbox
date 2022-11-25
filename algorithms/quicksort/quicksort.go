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

	// p := partition(slice, low, high)
	// p := partition2(slice, low, high)
	p := partition3(slice, low, high)
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

// partition2 is an alternative partitioning scheme wherein
// variables are compared to the pivot from both sides of the slice
func partition2[T sortable](slice []T, low, high int) int {
	pivot := slice[low]

	i := low
	j := high

	for i < j {
		for slice[i] <= pivot && i < high {
			i++
		}

		for slice[j] > pivot && j > low {
			j--
		}

		if i < j {
			slice[i], slice[j] = slice[j], slice[i]
		}
	}

	slice[j], slice[low] = slice[low], slice[j]

	return j
}

func partition3[T sortable](slice []T, low, high int) int {
	pivot := slice[low]

	i := low
	j := i

	for j < high {
		j++
		if slice[j] < pivot {
			i++
			slice[j], slice[i] = slice[i], slice[j]
		}
	}

	slice[low], slice[i] = slice[i], slice[low]

	return i
}
