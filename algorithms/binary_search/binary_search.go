package main

type sortable interface {
	uint8 | uint16 | uint32 | uint64 |
		int | int8 | int16 | int32 | int64 |
		float32 | float64 |
		string
}

const NotFound = -1

// BinarySearch recursively
func BinarySearch[T sortable](target T, orderedCollection []T) int {
	return binarySearch(target, orderedCollection, 0, len(orderedCollection)-1)
}

func binarySearch[T sortable](target T, orderedCollection []T, low, high int) int {
	if low > high {
		return NotFound
	}

	midpoint := (low + high) / 2
	if orderedCollection[midpoint] == target {
		return midpoint
	}

	if orderedCollection[midpoint] > target {
		return binarySearch(target, orderedCollection, low, midpoint-1)
	}

	return binarySearch(target, orderedCollection, midpoint+1, high)
}

// BinarySearch with for loop
func BinarySearch2[T sortable](target T, orderedCollection []T) int {
	low := 0
	high := len(orderedCollection) - 1

	for low <= high {
		midpoint := (low + high) / 2
		if orderedCollection[midpoint] == target {
			return midpoint
		}

		if orderedCollection[midpoint] < target {
			low = midpoint + 1
		} else {
			high = midpoint - 1
		}
	}

	return NotFound
}
