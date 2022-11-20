package main

// quicksort can sort a slice of integers
func quicksort(s []int, low int, high int) {
	if low >= high {
		return
	}

	p := partition(s, low, high)
	quicksort(s, low, p-1)
	quicksort(s, p+1, high)
}

// partition is responsible for ensuring some pivot index is
// put in its final position. It returns that index
func partition(s []int, low int, high int) int {
	pivot := s[high]
	i := low

	for j := low; j < high; j++ {
		if s[j] < pivot {
			s[i], s[j] = s[j], s[i]
			i++
		}
	}

	s[i], s[high] = s[high], s[i]

	return i
}
