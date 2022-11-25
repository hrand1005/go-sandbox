package main

type sortable interface {
	uint8 | uint16 | uint32 | uint64 |
		int | int8 | int16 | int32 | int64 |
		float32 | float64 |
		string
}

func mergesort[T sortable](slice []T) []T {
	if len(slice) < 2 {
		return slice
	}

	midpoint := len(slice) / 2
	first := mergesort(slice[:midpoint])
	second := mergesort(slice[midpoint:])

	return merge(first, second)
}

func merge[T sortable](slice1, slice2 []T) []T {
	i := 0
	j := 0
	result := make([]T, 0, len(slice1)+len(slice2))

	for i < len(slice1) && j < len(slice2) {
		if slice1[i] < slice2[j] {
			result = append(result, slice1[i])
			i++
		} else {
			result = append(result, slice2[j])
			j++
		}
	}

	if i == len(slice1) {
		return append(result, slice2[j:]...)
	}
	return append(result, slice1[i:]...)
}
